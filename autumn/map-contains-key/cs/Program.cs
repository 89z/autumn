using C = System.Console;
using G = System.Collections.Generic;

class Program {
   static void Main() {
      var m = new G.Dictionary<string, int>{{"month", 5}, {"day", 4}};
      var b = m.ContainsKey("day");
      C.WriteLine(b);
   }
}
