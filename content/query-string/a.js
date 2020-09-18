// example A
let s = 'one=odd&two=even';
let oA = new URLSearchParams(s);
// example B
let m = {one: 'odd', two: 'even'};
let sB = new URLSearchParams(m).toString();
// print
console.log(oA, sB);
