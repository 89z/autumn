using System;
using System.Collections;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      Array.Sort(a);
      Console.WriteLine(Serialize(a));
   }
}
