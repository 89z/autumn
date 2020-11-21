using System;

class Program {
   static void Main() {
      var n = 0x21 + 1_000;
      var n2 = 1e3;
      Console.WriteLine(n == 1033 && n2 == 1000);
   }
}
