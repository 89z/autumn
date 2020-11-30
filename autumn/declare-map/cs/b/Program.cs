using C = System.Console;
using G = System.Collections.Generic;

class Program {
   static void Main() {
      var m = new G.Dictionary<string, int>{{"month", 12}, {"day", 31}};
      C.WriteLine(m);
   }
}
