// example 1
function f(s) {
   return s.length;
}
// example 2
let f2 = function (s) {
   return s.length;
};
// example 3
let f3 = s => s.length;
// print
let n = f('May');
let n2 = f2('May');
let n3 = f3('May');
console.log(n == 3, n2 == 3, n3 == 3);
