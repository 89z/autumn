// example 1
let t = new Set(['May', 'June']);
let b = t.has('June');
// example 2
let t2 = {May: true, June: true};
let b2 = 'June' in t2;
// print
console.log(b, b2);
