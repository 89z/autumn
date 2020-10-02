// example 1
let s = 'one=odd&two=even';
let o1 = new URLSearchParams(s);
// example 2
let m = {one: 'odd', two: 'even'};
let s2 = new URLSearchParams(m).toString();
// print
console.log(o1, s2);
