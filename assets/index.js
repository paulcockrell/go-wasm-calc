const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});

// global calc function
window.calc = function(op) {
    let v1 = Number(document.getElementById("val1").value);
    let v2 = Number(document.getElementById("val2").value);
    if (!v1 || !v2) {
        console.log("You must specify both numbers");
        return;
    }
    if (op === '+') {
        add(v1, v2);
    } else if (op === '-') {
        sub(v1, v2);
    } else if (op === '*') {
        mul(v1, v2);
    } else if (op === '/') {
        div(v1, v2);
    }
}