let o = new Date;
// example 1
let n1 = o.getTime();
let n2 = Math.trunc(n1 / 1000);
// example 2
let n3 = o.valueOf();
let n4 = Math.trunc(n3 / 1000);
// print
console.log(n2, n4);
