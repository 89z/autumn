// example 1
let n1 = 10;
do {
   console.log(n1);
   n1++;
} while (n1 < 20);
// example 2
let n2 = 10;
while (true) {
   console.log(n2);
   if (n2 == 19) {
      break;
   }
   n2++;
}
// example 3
let n3 = 10;
for (;;) {
   console.log(n3);
   if (n3 == 19) {
      break;
   }
   n3++;
}
