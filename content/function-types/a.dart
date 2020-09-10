// example 1
num f(String s) {
   return s.length;
}
// example 2
num f2(String s) => s.length;
// example 3
var f3 = (String s) => s.length;
// print
void main() {
   var n = f('May');
   var n2 = f2('May');
   var n3 = f3('May');
   print([n == 3, n2 == 3, n3 == 3]);
}
