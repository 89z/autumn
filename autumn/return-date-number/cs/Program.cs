using System;

class Program {
   static void Main() {
      var o = DateTimeOffset.Now;
      var n = o.ToUnixTimeSeconds();
      Console.WriteLine(n);
   }
}
