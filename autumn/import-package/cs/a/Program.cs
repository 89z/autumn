using System.Collections.Generic;
using System.Text.Json;
using System;

class Program {
   static void Main() {
      Dictionary<string, int>[] a = {
         new Dictionary<string, int>{{"year", 2019}, {"month", 12}},
         new Dictionary<string, int>{{"year", 2018}, {"month", 11}}
      };
      var s = JsonSerializer.Serialize(a);
      Console.WriteLine(s);
   }
}
