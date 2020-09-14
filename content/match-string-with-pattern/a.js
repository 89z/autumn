let s= 'January';
// example 1
let b = /^J/.test(s);
// example 2
let b2 = /ja/i.test(s);
// example 3
let b3 = RegExp(/ja/, 'i').test(s);
// print
console.log(b, b2, b3);
