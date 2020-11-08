using System.Text.Json;
using System;

class Program {
   static void Main() {
      var s = "[\"sigma\", \"tau\"]";
      var a = JsonSerializer.Deserialize<string[]>(s);
      Console.WriteLine(a);
   }
}
