// example 1
function f1(s1) {
   return s1.length;
}
// example 2
let f2 = function (s1) {
   return s1.length;
};
// example 3
let f3 = s1 => s1.length;
// print
let n1 = f1('Sunday');
let n2 = f2('Sunday');
let n3 = f3('Sunday');
console.log(n1, n2, n3);
