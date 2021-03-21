void main() {
   var s = "100";
   { // example 1
      var n = int.parse(s);
      print(n == 100);
   }
   { // example 2
      var n = num.parse(s);
      print(n == 100);
   }
}
