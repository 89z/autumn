using System;
using System.Linq;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      // example 1
      var n = a.Aggregate(0, (n, s) => n + s.Length);
      // example 2
      var s = a.Aggregate((s, s2) => s + s2);
      // print
      Console.WriteLine(n == 7 && s == "MayJune");
   }
}
