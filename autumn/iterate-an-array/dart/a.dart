void main() {
   var a = ['west', 'east'];
   // example 1
   for (var s in a) {
      print(s);
   }
   // example 2
   var f = (String s) => print(s);
   a.forEach(f);
   // example 3
   for (var n = 0; n < a.length; n++) {
      var s = a[n];
      print([n, s]);
   }
}
