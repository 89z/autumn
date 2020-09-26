using System;
using System.Linq;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      var s = a.Aggregate((s, s1) => s + s1);
      Console.WriteLine(s == "MayJune");
   }
}
