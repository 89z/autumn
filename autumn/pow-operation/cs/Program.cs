using System;

class Program {
   static void Main() {
      // example 1
      var n1 = Math.Pow(10, 5);
      // example 2
      var n2 = Math.Pow(9, 0.5);
      // print
      Console.WriteLine(n1 == 1e5 && n2 == 3);
   }
}
