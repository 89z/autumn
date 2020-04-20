let s1 = '2019-12-31';
// example 1
let n1 = s1.search(/\d/);
let b1 = n1 != -1;
// example 2
let o1 = s1.match(/\d/);
let b2 = o1 != null;
// example 3
let o2 = s1.match(/-../);
let a1 = Array.from(o2);
// example 4
let o3 = s1.match(/-../g);
let a2 = Array.from(o3);
// example 5
let o4 = s1.match(/-(..)-(..)/);
let a3 = Array.from(o4);
// print
console.log(b1, b2, a1, a2, a3);
