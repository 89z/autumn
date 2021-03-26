void main() {
   { // num
      var s = "100";
      var n = num.parse(s);
      print(n == 100);
   }
   { // int
      var s = "100";
      var n = int.parse(s);
      print(n == 100);
   }
   { // double
      var s = "99.9";
      var n = double.parse(s);
      print(n == 99.9);
   }
}
