let a = ['May', 'June'];
// example 1
let s = a.reduce((s, s1) => s + s1);
// example 2
let n = a.reduce((n, s) => n + s.length, 0);
// print
console.log(s == 'MayJune', n == 7);
