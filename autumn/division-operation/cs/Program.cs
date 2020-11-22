using System;

class Program {
   static void Main() {
      var (f, i) = (7.5, 7);
      // example 1
      var n1 = f / 2;
      // example 2
      var n2 = (int) f / 2;
      // example 3
      var n3 = i / 2;
      // example 4
      var n4 = (double) i / 2;
      // print
      Console.WriteLine(n1 == 3.75 && n2 == 3 && n3 == 3 && n4 == 3.5);
   }
}
