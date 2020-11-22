using System;

class Program {
   static void Main() {
      var (f, i) = (7.5, 7);
      // example 1
      var n1 = f % 2;
      // example 2
      var n2 = i % 2;
      // print
      Console.WriteLine(n1 == 1.5 && n2 == 1);
   }
}
