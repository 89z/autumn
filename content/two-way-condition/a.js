let n = 10;
// example 1
let s1;
if (n > 0) {
   s1 = '+';
} else if (n < 0) {
   s1 = '-';
} else {
   s1 = 'zero';
}
// example 2
let s2 = n > 0 ? '+' : '-';
// print
console.log(s1 == '+', s2 == '+');
