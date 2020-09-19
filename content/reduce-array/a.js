let a = ['May', 'June'];
// example 1
let s1 = a.reduce((s, s1) => s + s1);
// example 2
let f = (s, s2) => s + s2;
let s2 = a.reduce(f)
// print
console.log(s1 == 'MayJune', s2 == 'MayJune');
