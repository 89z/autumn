using C = System.Console;
using G = System.Collections.Generic;
using J = Newtonsoft.Json.JsonConvert;

class Program {
   static void Main() {
      var m = new G.Dictionary<string, int>{};
      m.Add("day", 31);
      C.WriteLine(J.SerializeObject(m));
   }
}
