function time(...a) {
   if (a.length == 1) {
      return new Date(a[0] * 1000);
   }
   a[1]--;
   return new Date(...a);
}

let t;
// example 1
t = time(0x5555_5555);
console.log(t);
// example 2
t = time(2020, 12);
console.log(t);
// example 3
t = time(2020, 12, 31);
console.log(t);
// example 4
t = time(2020, 12, 31, 23);
console.log(t);
// example 5
t = time(2020, 12, 31, 23, 59);
console.log(t);
// example 6
t = time(2020, 12, 31, 23, 59, 59);
console.log(t);
