let a1 = ['sun', 'mon', 'tue'];
// string from front
let s1 = a1[0];
// string from back
let s2 = a1[a1.length - 1];
// array from front
let a2 = a1.slice(0, 2);
// array from back
let a3 = a1.slice(-2);
// print
console.log(s1, s2, a2, a3);
