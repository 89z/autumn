let a = ['May', 'June'];
// example 1
let s = a.reduce((sa, sc) => sa + sc);
// example 2
let f = (sa, sc) => sa + sc;
let s2 = a.reduce(f);
// print
console.log(s == 'MayJune', s2 == 'MayJune');
