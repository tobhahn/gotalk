function insertTabs(n) {
    // find the selection start and end
    var cont  = document.getElementById("edit");
    var start = cont.selectionStart;
    var end   = cont.selectionEnd;
    // split the textarea content into two, and insert n tabs
    var v = cont.value;
    var u = v.substr(0, start);
    for (var i=0; i<n; i++) {
        u += "\t";
    }
    u += v.substr(end);
    // set revised content
    cont.value = u;
    // reset caret position after inserted tabs
    cont.selectionStart = start+n;
    cont.selectionEnd = start+n;
}

function autoindent(el) {
    var curpos = el.selectionStart;
    var tabs = 0;
    while (curpos > 0) {
        curpos--;
        if (el.value[curpos] == "\t") {
            tabs++;
        } else if (tabs > 0 || el.value[curpos] == "\n") {
            break;
        }
    }
    setTimeout(function() {
        insertTabs(tabs);
    }, 1);
}

function preventDefault(e) {
    if (e.preventDefault) {
        e.preventDefault();
    } else {
        e.cancelBubble = true;
    }
}

function keyHandler(event) {
    var e = window.event || event;
    if (e.keyCode == 9) { // tab
        insertTabs(1);
        preventDefault(e);
        return false;
    }
    if (e.keyCode == 13) { // enter
        if (e.shiftKey) { // +shift
            compile(e.target);
            preventDefault(e);
            return false;
        } else {
            autoindent(e.target);
        }
    }
    return true;
}

var xmlreq;

function autocompile() {
    if(!document.getElementById("autocompile").checked) {
        return;
    }
    compile();
}

function compile() {
    var prog = document.getElementById("edit").value;
    var req = new XMLHttpRequest();
    xmlreq = req;
    req.onreadystatechange = compileUpdate;
    req.open("GET", "/compile?q=" + encodeURIComponent(prog), true);
    req.setRequestHeader("Content-Type", "text/plain; charset=utf-8");
    req.send();
    $("#output").addClass("hide");
    $("#errors").addClass("hide");
}

function compileUpdate() {
    var req = xmlreq;
    if(!req || req.readyState != 4) {
        return;
    }
    if(req.status == 200) {
        document.getElementById("output").innerHTML = req.responseText;
        document.getElementById("errors").innerHTML = "";
        $("#output").removeClass("hide");

    } else {
        document.getElementById("errors").innerHTML = req.responseText;
        document.getElementById("output").innerHTML = "";
        $("#errors").removeClass("hide");
    }
}
