// example 1
let c1 = new Set(['sun', 'mon']);
let b1 = c1.has('mon');
// example 2
let c2 = {sun: true, mon: true};
let b2 = 'mon' in c2;
// print
console.log(b1, b2);
