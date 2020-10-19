using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "M a r c h";
      // example 1
      var a1 = s.Split(" ");
      // example 2
      var a2 = s.Split();
      // print
      Console.WriteLine(Serialize(a1) + Serialize(a2));
   }
}
