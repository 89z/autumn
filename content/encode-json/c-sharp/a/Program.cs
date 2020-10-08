using System.Collections.Generic;
using System.Text.Json;
using System;

class Program {
   static void Main() {
      // example 1
      string[] a = {"May", "June"};
      var s1 = JsonSerializer.Serialize(a);
      // example 2
      var m = new Dictionary<string, int> {
         {"year", 2019}, {"month", 12}
      };
      var s2 = JsonSerializer.Serialize(m);
      // print
      Console.WriteLine("{0} {1}", s1, s2);
   }
}
