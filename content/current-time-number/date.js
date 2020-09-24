let o = new Date;
// example 1
let n = o.getTime();
let n1 = Math.trunc(n / 1000);
// example 2
let n2a = o.valueOf();
let n2 = Math.trunc(n2a / 1000);
// print
console.log(n1, n2);
