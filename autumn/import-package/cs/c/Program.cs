using Dict = System.Collections.Generic.Dictionary<string, int>;
using System.Text.Json;
using System;

class Program {
   static void Main() {
      Dict[] a = {
         new Dict {{"year", 2019}, {"month", 12}},
         new Dict {{"year", 2018}, {"month", 11}},
      };
      var s = JsonSerializer.Serialize(a);
      Console.WriteLine(s);
   }
}
