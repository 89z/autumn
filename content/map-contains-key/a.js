let m = {'Sunday': 10};
// example 1
let b1 = 'Sunday' in m;
// example 2
let b2 = m.hasOwnProperty('Sunday');
// example 3
let b3 = Reflect.has(m, 'Sunday');
// print
console.log(b1, b2, b3);
