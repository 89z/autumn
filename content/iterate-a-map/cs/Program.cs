using System;
using System.Collections.Generic;

class Program {
   static void Main() {
      var m = new Dictionary<string, int>();
      m.Add("year", 2019);
      foreach (KeyValuePair<string, int> o in m) {
         Console.WriteLine("{0} {1}", o.Key, o.Value);
      }
   }
}
