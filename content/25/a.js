let s = '2.9';
// example 1
let n = Number(s);
// example 2
let n2 = Number.parseFloat(s);
// example 3
let n3 = parseFloat(s);
// example 4
let n4 = +(s);
// print
console.log(n === 2.9, n2 === 2.9, n3 === 2.9, n4 === 2.9);
