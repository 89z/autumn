let s = 'Wednesday';
// example 1
let a = s.match(/e./);
// example 2
let a2 = s.match(/e./g);
// example 3
let o = s.matchAll(/e./g);
let a3 = Array.from(o);
// example 4
let a4 = s.match(/e(..)/);
// example 5
let o2 = s.matchAll(/e(..)/g);
let a5 = Array.from(o2);
// print
console.log(a, a2, a3, a4, a5);
