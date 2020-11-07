using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      Console.WriteLine(Serialize(a));
   }
}
