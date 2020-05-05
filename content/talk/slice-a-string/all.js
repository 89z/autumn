let s1 = 'Sunday';
// one from start
let s2 = s1.slice(0, 1);
// two from start
let s3 = s1.slice(0, 2);
// one from middle
let s4 = s1.slice(1, 2);
// from middle to end
let s5 = s1.slice(1);
// one from end
let s6 = s1.slice(-1);
// two from end
let s7 = s1.slice(-2);
// print
console.log(
   s2 == 'S', s3 == 'Su', s4 == 'u', s5 == 'unday', s6 == 'y', s7 == 'ay'
);
