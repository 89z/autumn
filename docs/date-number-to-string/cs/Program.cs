using System;

class Program {
   static void Main() {
      var n = 1577858399;
      var o = DateTimeOffset.FromUnixTimeSeconds(n);
      var s = o.LocalDateTime.ToString();
      Console.WriteLine(s == "2019-12-31 11:59:59 PM");
   }
}
