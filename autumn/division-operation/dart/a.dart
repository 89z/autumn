void main() {
   { // natural float
      var n = 7.0 / 2;
      print(n == 3.5);
   }
   { // natural float
      var n = 7 / 2;
      print(n == 3.5);
   }
   { // force int
      var n = 7.0 ~/ 2;
      print(n == 3);
   }
   { // force int
      var n = 7 ~/ 2;
      print(n == 3);
   }
}
