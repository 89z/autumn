void main() {
   var f = 7.5, i = 7;
   { // example 1
      var n = f / 2;
      print(n == 3.75);
   }
   { // example 2
      var n = f ~/ 2;
      print(n == 3);
   }
   { // example 3
      var n = i ~/ 2;
      print(n == 3);
   }
   { // example 4
      var n = i / 2;
      print(n == 3.5);
   }
}
