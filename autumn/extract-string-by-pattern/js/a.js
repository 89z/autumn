let s = 'January';
// example 1
let a1 = s.match(/a./);
// example 2
let a2 = s.match(/a./g);
// example 3
let o = s.matchAll(/a./g);
let a3 = Array.from(o);
// example 4
let a4 = s.match(/a(..)/);
// example 5
let o5 = s.matchAll(/a(..)/g);
let a5 = Array.from(o5);
// print
console.log(a1, a2, a3, a4, a5);
