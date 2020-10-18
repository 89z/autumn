using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = {"J", "u", "n", "e"};
      var a2 = a[2 ..];
      Console.WriteLine(Serialize(a2));
   }
}
