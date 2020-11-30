void main() {
   // example 1
   var s1 = 'zero"one\\two';
   // example 2
   var s2 = "zero\"one\\two";
   // example 3
   var s3 = r'zero"one\two';
   // example 4
   var s4 = r"one\two";
   // print
   print(s1 + ',' + s2 + ',' + s3 + ',' + s4);
}
