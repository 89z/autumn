let m = {'year': 2019};
// example 1
let b1 = 'year' in m;
// example 2
let b2 = m.hasOwnProperty('year');
// example 3
let b3 = Reflect.has(m, 'year');
// print
console.log(b1, b2, b3);
