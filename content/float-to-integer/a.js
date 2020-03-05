let n1 = 1.9;
// truncate
let n2 = Math.trunc(n1);
// truncate: fails with 10e9
let n3 = n1 | 0;
// truncate: fails with 10e9
let n4 = ~~n1;
// round
let n5 = Math.round(n1);
// print
console.log(n2, n3, n4, n5);
