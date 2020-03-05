let m1 = {'sun': 10, 'mon': 11};
// example 1
let b1 = 'mon' in m1;
// example 2
let b2 = m1.hasOwnProperty('mon');
// example 3
let b3 = Reflect.has(m1, 'mon');
// print
console.log(b1, b2, b3);
