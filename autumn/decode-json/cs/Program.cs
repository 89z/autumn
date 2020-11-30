using C = System.Console;
using G = System.Collections.Generic;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "{\"month\": 12, \"day\": 31}";
      var m = J.Deserialize<G.Dictionary<string, int>>(s);
      C.WriteLine(m["day"] == 31);
   }
}
