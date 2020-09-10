let n = 1;
// example 1
let s;
if (n > 0) {
   s = '+';
} else if (n < 0) {
   s = '-';
} else {
   s = 'zero';
}
// example 2
let s2 = n > 0 ? '+' : '-';
// print
console.log(s == '+', s2 == '+');
