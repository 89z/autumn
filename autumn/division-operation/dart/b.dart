void main() {
   var f = 7.5, i = 7;
   { // example 1
      var n = (f / 2).toInt();
      print(n == 3);
   }
   { // example 2
      var n = (i / 2).toInt();
      print(n == 3);
   }
}
