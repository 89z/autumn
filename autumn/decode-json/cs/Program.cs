using C = System.Console;
using G = System.Collections.Generic;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "{\"year\": 2019, \"month\": 12}";
      var m = J.Deserialize<G.Dictionary<string, int>>(s);
      C.WriteLine(m["year"]);
   }
}
