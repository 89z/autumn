let s1 = '2019-12-31';
// example 1
let b1 = /\d/.test(s1);
// example 2
let o1 = /\d/.exec(s1);
let b2 = o1 != null;
// example 3
let o2 = /-../.exec(s1);
let a1 = Array.from(o2);
// example 4
let o3 = /-(..)-(..)/.exec(s1);
let a2 = Array.from(o3);
// print
console.log(b1, b2, a1, a2);
