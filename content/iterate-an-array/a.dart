main() {
   var a = ['Sunday', 'Monday'];
   // example 1
   for (var s in a) {
      print(s);
   }
   // example 2
   a.forEach((s) => print(s));
   // example 3
   a.forEach(print);
}
