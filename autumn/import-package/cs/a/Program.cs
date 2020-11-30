using System.Collections.Generic;
using System.Text.Json;
using System;

class Program {
   static void Main() {
      Dictionary<string, int>[] a = {
         new Dictionary<string, int>{{"month", 12}, {"day", 31}},
         new Dictionary<string, int>{{"month", 11}, {"day", 30}},
      };
      Console.WriteLine(JsonSerializer.Serialize(a));
   }
}
