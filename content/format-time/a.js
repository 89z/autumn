let n1 = 86400;
let o1 = new Date(n1 * 1000);
// example 1
let s1 = o1.toString();
// example 2
let s2 = o1.toISOString();
// example 3
let s3 = o1.toDateString();
// example 4
let s4 = o1.toLocaleString();
// print
console.log(s1, s2, s3, s4);
