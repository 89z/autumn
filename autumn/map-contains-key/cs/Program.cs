using C = System.Console;
using G = System.Collections.Generic;

class Program {
   static void Main() {
      var m = new G.Dictionary<string, int>{{"year", 2019}, {"month", 12}};
      var b = m.ContainsKey("year");
      C.WriteLine(b);
   }
}
