// example 1
let s1 = 'one=odd&two=even';
let o1 = new URLSearchParams(s1);
// example 2
let m1 = {one: 'odd', two: 'even'};
let s2 = new URLSearchParams(m1).toString();
// print
console.log(o1, s2);
