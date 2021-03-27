void main() {
   var a = [0, 10, 20, 30, 40];
   { // example 1
      var b = a.skip(2).take(2);
      print(b);
   }
   { // example 2
      var b = a.skip(2);
      print(b);
   }
}
