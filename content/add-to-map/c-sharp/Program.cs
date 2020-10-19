using System.Collections.Generic;
using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var m = new Dictionary<string, int>();
      m.Add("year", 2019);
      Console.WriteLine(Serialize(m));
   }
}
