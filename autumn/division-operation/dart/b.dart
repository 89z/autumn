void main() {
   { // force int
      var n = (7.0 / 2).toInt();
      print(n == 3);
   }
   { // force int
      var n = (7 / 2).toInt();
      print(n == 3);
   }
}
