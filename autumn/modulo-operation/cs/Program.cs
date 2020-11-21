using System;

class Program {
   static void Main() {
      // example 1
      var n1 = 46 % 10;
      // example 2
      var n2 = 46 / 10;
      // example 3
      var n3 = 46 / 10d;
      // print
      Console.WriteLine(n1 == 6 && n2 == 4 && n3 == 4.6);
   }
}
