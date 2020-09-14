let s = '9.9';
// example 1
let n = Number(s);
// example 2
let n2 = Number.parseFloat(s);
// example 3
let n3 = parseFloat(s);
// example 4
let n4 = +(s);
// print
console.log(n === 9.9, n2 === 9.9, n3 === 9.9, n4 === 9.9);
