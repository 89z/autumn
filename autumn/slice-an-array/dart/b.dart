void main() {
   var a = ['M', 'a', 'r', 'c', 'h'];
   { // example 1
      var b = a.skip(2).take(2);
      print(b);
   }
   { // example 2
      var b = a.skip(2);
      print(b);
   }
}
