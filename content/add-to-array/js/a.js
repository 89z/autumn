let a = [10, 11];
// example 1
a.push(12);
// example 2
a = a.concat(13);
// example 3
a = a.concat(14, 15);
// example 4
let a4 = [16, 17];
a = a.concat(a4);
// example 5
a = [...a, 18];
// print
console.log(a);
