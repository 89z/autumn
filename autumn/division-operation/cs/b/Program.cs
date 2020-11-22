using System;

class Program {
   static void Main() {
      var n = 7.5;
      // example 1
      var n1 = n / 2;
      // example 2
      var n2 = (int) n / 2;
      // print
      Console.WriteLine(n1 == 3.75 && n2 == 3);
   }
}
