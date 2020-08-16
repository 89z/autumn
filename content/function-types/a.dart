// example 1
f1(s) {
   return s.length;
}
// example 2
f2(s) => s.length;
// example 3
var f3 = (s) => s.length;
// print
main() {
   var s = 'Sunday';
   print(f1(s));
   print(f2(s));
   print(f3(s));
}
