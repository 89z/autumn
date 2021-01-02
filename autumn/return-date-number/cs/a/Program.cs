using System;

class Program {
   static void Main() {
      var o = DateTimeOffset.Now;
      var n = o.ToUnixTimeMilliseconds();
      Console.WriteLine(n);
   }
}
