void main() {
   var a = ['May', 'June'];
   // example 1
   for (var s in a) {
      print(s);
   }
   // example 2
   var f = (s) => print(s);
   a.forEach(f);
   // example 3
   a.forEach(print);
}
