let s1 = 'Sunday';
// example 1
let b1 = /^Su/.test(s1);
// example 2
let b2 = /un./.test(s1);
// example 3
let b3 = /ay$/.test(s1);
// example 4
let b4 = /su/i.test(s1);
// example 5
let b5 = RegExp(/su/, 'i').test(s1);
// print
console.log(b1, b2, b3, b4, b5);
