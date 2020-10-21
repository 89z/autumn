using System;
using System.Collections.Generic;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var a = new List<string>();
      a.Add("May");
      Console.WriteLine(Serialize(a));
   }
}
