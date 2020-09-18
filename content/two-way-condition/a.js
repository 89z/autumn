let n = 10;
// example A
let sA;
if (n > 0) {
   sA = '+';
} else if (n < 0) {
   sA = '-';
} else {
   sA = 'zero';
}
// example B
let sB = n > 0 ? '+' : '-';
// print
console.log(sA == '+', sB == '+');
