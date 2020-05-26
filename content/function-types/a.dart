// example 1
f1(s) {
   return s.length;
}
// example 2
f2(s) => s.length;
// print
main() {
   var n1 = f1('Sunday');
   print(n1);
   var n2 = f2('Sunday');
   print(n2);
}
