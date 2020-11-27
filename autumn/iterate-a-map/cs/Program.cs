using C = System.Console;
using G = System.Collections.Generic;

class Program {
   static void Main() {
      var m = new G.Dictionary<string, int>{{"year", 2019}, {"month", 12}};
      foreach (G.KeyValuePair<string, int> o in m) {
         C.WriteLine("{0} {1}", o.Key, o.Value);
      }
   }
}
