// example 1
let a1 = ['Sun', 'Mon'];
let t1 = new Set(a1);
let b1 = t1.has('Mon');
// example 2
let t2 = {Sun: true, Mon: true};
let b2 = 'Mon' in t2;
// print
console.log(b1, b2);
