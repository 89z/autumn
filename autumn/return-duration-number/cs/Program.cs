using System;

class Program {
   static double SinceHours(string s) {
      var t = DateTime.Parse(s);
      return (DateTime.Now - t).TotalHours;
   }
   static void Main() {
      var n = SinceHours("2020-12-31");
      Console.WriteLine(n);
   }
}
