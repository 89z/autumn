using C = System.Console;
using G = System.Collections.Generic;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var a = new G.List<int>{9, 8};
      a.Add(7);
      C.WriteLine(J.Serialize(a));
   }
}
