let s = '📗';
// example 1
let n = Array.from(s).length;
// example 2
let n2 = s.length;
// example 3
let o = new TextEncoder;
let n3 = o.encode(s).length;
// print
console.log(n == 1, n2 == 2, n3 == 4);
