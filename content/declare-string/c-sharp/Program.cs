using System;

class Program {
   static void Main() {
      // example 1
      var s1 = "one\\two";
      // example 2
      var s2 = @"one\two";
      // print
      Console.WriteLine(s1 + "," + s2);
   }
}
