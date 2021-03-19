let te = new TextEncoder;
let a = te.encode('March');
let s = String.fromCharCode(...a);
console.log(s == 'March')
