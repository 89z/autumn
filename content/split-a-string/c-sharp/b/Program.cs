using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "March";
      var a = s.ToCharArray();
      Console.WriteLine(Serialize(a));
   }
}
