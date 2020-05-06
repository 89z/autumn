let s1 = 'package';
// example 1
let a1 = s1.match(/p./);
// example 2
let a2 = s1.match(/a./g);
// example 3
let a3 = s1.match(/p(..)/);
// print
console.log(a1, a2, a3);
