let s1 = 'Sunday';
// example 1
let b1 = s1.startsWith('Su');
// example 2
let b2 = /^Su/.test(s1);
// example 3
let b3 = s1.includes('un');
// example 4
let b4 = /un./.test(s1);
// example 5
let b5 = s1.endsWith('ay');
// example 6
let b6 = /ay$/.test(s1);
// print
console.log(b1, b2, b3, b4, b5, b6);
