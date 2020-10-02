// example 1
let t = new Set(['May', 'June']);
let b1 = t.has('May');
// example 2
let t2 = {May: true, June: true};
let b2 = 'June' in t2;
// print
console.log(b1, b2);
