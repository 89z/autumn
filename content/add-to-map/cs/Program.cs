using System;
using System.Collections.Generic;

class Program {
   static void Main() {
      var m = new Dictionary<string, int>();
      m.Add("year", 2019);
      m.Add("month", 12);
      foreach (var o in m) {
         Console.WriteLine("{0} {1}", o.Key, o.Value);
      }
   }
}
