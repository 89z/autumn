using System.Collections.Generic;
using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var t = new HashSet<string>{};
      t.Add("May");
      Console.WriteLine(Serialize(t));
   }
}
