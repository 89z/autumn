let n1 = 0;
// example 1
let s1;
if (n1 > 0) {
   s1 = 'positive';
} else if (n1 < 0) {
   s1 = 'negative';
} else {
   s1 = 'zero';
}
// example 2
let s2 = n1 < 0 ? 'negative' : 'positive';
// print
console.log(s1, s2);
