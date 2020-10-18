using System.Linq;
using System;
using static System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = {"M", "a", "r", "c", "h"};
      // example 1
      var a1 = a.Skip(2).Take(2);
      // example 2
      var a2 = a.Skip(2);
      // print
      Console.WriteLine(Serialize(a1) + Serialize(a2));
   }
}
