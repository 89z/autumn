let m1 = {'Sun': 10, 'Mon': 11};
// example 1
let b1 = 'Mon' in m1;
// example 2
let b2 = m1.hasOwnProperty('Mon');
// example 3
let b3 = Reflect.has(m1, 'Mon');
// print
console.log(b1, b2, b3);
