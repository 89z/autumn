// example 1
let c1 = new Set(['Sun', 'Mon']);
let b1 = c1.has('Mon');
// example 2
let c2 = {Sun: true, Mon: true};
let b2 = 'Mon' in c2;
// print
console.log(b1, b2);
