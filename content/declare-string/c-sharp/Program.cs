using System;

class Program {
   static void Main() {
      // example 1
      string s1 = default;
      // example 2
      string s2 = "one\\two";
      // example 3
      var s3 = "one\\two";
      // example 4
      var s4 = @"one\two";
      // print
      Console.WriteLine("{0},{1},{2},{3}", s1, s2, s3, s4);
   }
}
