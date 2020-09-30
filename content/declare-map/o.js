let m = {year: 2019};
// example 1
m['month'] = 12;
// example 2
m.day = 31;
// example 3
let m3 = {hour: 23};
m = Object.assign(m, m3);
// example 4
let m4 = {minute: 59};
m = {...m, ...m4};
// print
console.log(m);
