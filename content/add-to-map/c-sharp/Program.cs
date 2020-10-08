using System.Collections.Generic;
using System.Text.Json;
using System;

class Program {
   static void Main() {
      var m = new Dictionary<string, int>();
      m.Add("year", 2019);
      m.Add("month", 12);
      Console.WriteLine(JsonSerializer.Serialize(m));
   }
}
