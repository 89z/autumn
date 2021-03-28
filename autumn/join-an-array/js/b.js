let t = new TextEncoder;
let a = t.encode('north');
let s = String.fromCharCode(...a);
console.log(s == 'north')
