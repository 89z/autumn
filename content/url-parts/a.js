// example 1
let s1 = 'http://sun.com/mon?tue=10';
let o1 = new URL(s1);
// example 2
let s2 = 'http://sun.com';
let s3 = 'mon?tue=10';
let o2 = new URL(s3, s2);
// print
console.log(o1, o2);
