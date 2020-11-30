using C = System.Console;
using G = System.Collections.Generic;

class Program {
   static void Main() {
      var m = new G.Dictionary<string, int>{{"month": 12}, {"day": 31}};
      var n = m["day"];
      C.WriteLine(n == 31);
   }
}
