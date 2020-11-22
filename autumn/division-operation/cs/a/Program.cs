using System;

class Program {
   static void Main() {
      var n = 7;
      // example 1
      var n1 = n / 2;
      // example 2
      var n2 = (double) n / 2;
      // print
      Console.WriteLine(n1 == 3 && n2 == 3.5);
   }
}
