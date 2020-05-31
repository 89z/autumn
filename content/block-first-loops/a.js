let n;
// example 1
n = 10;
do {
   console.log(n);
   n++;
} while (n < 20);
// example 2
n = 10;
while (true) {
   console.log(n);
   if (n == 19) {
      break;
   }
   n++;
}
// example 3
n = 10;
for (;;) {
   console.log(n);
   if (n == 19) {
      break;
   }
   n++;
}
