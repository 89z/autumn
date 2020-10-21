using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var a = new[] {"May", "June"};
      Console.WriteLine(Serialize(a));
   }
}
