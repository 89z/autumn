let n = 10;
// example 1
let s1;
if (n > 12) {
   s1 = 'Tue';
} else if (n > 11) {
   s1 = 'Mon';
} else {
   s1 = 'Sun';
}
// example 2
let s2 = n > 11 ? 'Mon' : 'Sun';
// print
console.log(s1, s2);
