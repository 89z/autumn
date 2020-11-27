using C = System.Console;
using G = System.Collections.Generic;

class Program {
   static void Main() {
      var m = new G.Dictionary<string, int>{{"year", 2019}, {"month", 12}};
      var n = m["year"];
      C.WriteLine(n == 2019);
   }
}
