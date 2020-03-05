let n1 = 10;
// example 1
let n2 = n1 == 0 ? 20 : n1;
// example 2
let n3;
if (n1 > 30) {
   n3 = 30;
} else if (n1 > 20) {
   n3 = 20;
} else {
   n3 = n1;
}
// print
console.log(n2, n3);
