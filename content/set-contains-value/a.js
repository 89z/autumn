// example 1
let t1 = {Sunday: true, Monday: true};
let b1 = 'Sunday' in t1;
// example 2
let t2 = new Set(['Sunday', 'Monday']);
let b2 = t2.has('Sunday');
// print
console.log(b1, b2);
