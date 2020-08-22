let s1 = '📕';
// example 1
let a1 = Array.from(s1);
let n1 = a1.length;
// example 2
let n2 = s1.length;
// example 3
let o1 = new TextEncoder;
let a2 = o1.encode(s1);
let n3 = a2.length;
// print
console.log(n1 == 1, n2 == 2, n3 == 4);
