using System;
using System.Collections.Generic;

class Program {
   static void Main() {
      var m = new Dictionary<string, int>();
      m.Add("year", 2019);
      var n = m["year"];
      Console.WriteLine(n == 2019);
   }
}
