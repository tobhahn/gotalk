package gotalk

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Compile is an HTTP handler that reads Go source code from the request,
// runs the program (returning any errors),
// and sends the program's output as the HTTP response.
func compile(w http.ResponseWriter, req *http.Request) {
	out, err := compileRequest(req)
	if err != nil {
		error_(w, out, err)
		return
	}

	// write the output of x as the http response
	w.Write(out)
}

var (
	commentRe = regexp.MustCompile(`(?m)^#.*\n`)
	tmpdir    string
	// a source of numbers, for naming temporary files
	uniq = make(chan int)
	// timeout for compiling and running
	cmdTimeout = 3 * time.Second
)

func init() {
	// find real temporary directory (for rewriting filename in output)
	var err error
	tmpdir, err = filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		log.Fatal(err)
	}

	// source of unique numbers
	go func() {
		for i := 0; ; i++ {
			uniq <- i
		}
	}()
}

func compileRequest(req *http.Request) (out []byte, err error) {
	// x is the base name for .go, .6, executable files
	x := filepath.Join(tmpdir, "compile"+strconv.Itoa(<-uniq))
	src := x + ".go"
	bin := x
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	// rewrite filename in error output
	defer func() {
		if err != nil {
			// drop messages from the go tool like '# _/compile0'
			out = commentRe.ReplaceAll(out, nil)
		}
		out = bytes.Replace(out, []byte(src+":"), []byte("main.go:"), -1)
	}()

	// write body to x.go
	body := new(bytes.Buffer)
	if _, err = body.ReadFrom(strings.NewReader(req.URL.Query().Get("q"))); err != nil {
		return
	}
	defer os.Remove(src)
	if err = ioutil.WriteFile(src, body.Bytes(), 0666); err != nil {
		return
	}

	// build x.go, creating x
	dir, file := filepath.Split(src)
	out, err = run(dir, "go", "build", "-o", bin, file)
	defer os.Remove(bin)
	if err != nil {
		return
	}

	// run x
	return run("", bin)
}

// error writes compile, link, or runtime errors to the HTTP connection.
// The JavaScript interface uses the 404 status code to identify the error.
func error_(w http.ResponseWriter, out []byte, err error) {
	w.WriteHeader(404)
	if out != nil {
		w.Write(out)
	} else {
		w.Write([]byte(err.Error()))
	}
}

// run executes the specified command and returns its output and an error.
func run(dir string, args ...string) (output []byte, err error) {
	var buf bytes.Buffer
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = &buf
	cmd.Stderr = cmd.Stdout

	result := make(chan error)
	go func() {
		result <- cmd.Run()
	}()

	select {
	case err = <-result: // Command returned
	case <-time.After(cmdTimeout): // Timeout
		if err = cmd.Process.Kill(); err == nil {
			err = errors.New("Timeout")
			buf.Write([]byte(err.Error()))
		}
		<-result // let goroutine finish
	}
	return buf.Bytes(), err
}
