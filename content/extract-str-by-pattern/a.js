let a1, o1, s1 = 'Wednesday';
// example 1
a1 = s1.match(/e./);
console.log(a1);
// example 2
a1 = s1.match(/e./g);
console.log(a1);
// example 3
o1 = s1.matchAll(/e./g);
a1 = Array.from(o1);
console.log(a1);
// example 4
a1 = s1.match(/e(..)/);
console.log(a1);
// example 5
o1 = s1.matchAll(/e(..)/g);
a1 = Array.from(o1);
console.log(a1);
