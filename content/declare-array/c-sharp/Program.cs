using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      // example 1
      string[] a1 = {"May", "June"};
      // example 2
      var a2 = new[] {"May", "June"};
      // print
      Console.WriteLine(Serialize(a1) + Serialize(a2));
   }
}

