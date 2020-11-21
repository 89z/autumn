using System;

class Program {
   static void Main() {
      var n = 0x21 + 1_000 + 1e3;
      Console.WriteLine(n == 2033);
   }
}
