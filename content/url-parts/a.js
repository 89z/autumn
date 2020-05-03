// example 1
let s1 = 'https://example.com/one?two=even';
let o1 = new URL(s1);
// example 2
let s2 = 'https://example.com';
let s3 = 'one?two=even';
let o2 = new URL(s3, s2);
// print
console.log(o1, o2);
