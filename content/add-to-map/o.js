let m = {year: 2020};
// example 1
m['month'] = 9;
// example 2
m.day = 7;
// example 3
let m2 = {hour: 11};
m = Object.assign(m, m2);
// example 4
let m3 = {minute: 4};
m = {...m, ...m3};
// print
console.log(m);
