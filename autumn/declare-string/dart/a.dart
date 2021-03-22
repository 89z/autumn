void main() {
   { // example 1
      var s = 'zero"one\\two';
      print(s);
   }
   { // example 2
      var s = "zero\"one\\two";
      print(s);
   }
   { // example 3
      var s = r'zero"one\two';
      print(s);
   }
   { // example 4
      var s = r"one\two";
      print(s);
   }
}
