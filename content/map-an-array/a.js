let a1 = [16, 25];
// example 1
let a2 = a1.map(n1 => Math.sqrt(n1));
// example 2
let a3 = a1.map(Math.sqrt);
// example 3
let a4 = Array.from(a1, n1 => Math.sqrt(n1));
// print
console.log(a2, a3, a4);
