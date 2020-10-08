using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      var s = Serialize(a);
      Console.WriteLine(s);
   }
}
