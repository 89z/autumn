let s = 'Wednesday';
// example 1
let b1 = /^W/.test(s);
// example 2
let b2 = /we/i.test(s);
// example 3
let b3 = RegExp(/we/, 'i').test(s);
// print
console.log(b1, b2, b3);
