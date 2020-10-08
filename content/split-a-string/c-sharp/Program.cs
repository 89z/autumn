using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "May June";
      var a = s.Split();
      Console.WriteLine(Serialize(a));
   }
}
