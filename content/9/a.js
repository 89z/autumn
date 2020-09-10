let m = {'year': 2020};
// example 1
let b = 'year' in m;
// example 2
let b2 = m.hasOwnProperty('year');
// example 3
let b3 = Reflect.has(m, 'year');
// print
console.log(b, b2, b3);
