using System;

class Program {
   static void Main() {
      // example 1
      var s1 = @"sigma\tau";
      // example 2
      var s2 = @"[""sigma"", ""tau""]";
      // print
      Console.WriteLine(s1 + s2);
   }
}
