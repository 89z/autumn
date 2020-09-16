// example 1
function f(n, n2) {
   return n > n2;
}
// example 2
let f2 = function (n, n2) {
   return n > n2;
};
// example 3
let f3 = (n, n2) => n > n2;
// print
let b = f(9, 8);
let b2 = f2(9, 8);
let b3 = f3(9, 8);
console.log(b, b2, b3);
