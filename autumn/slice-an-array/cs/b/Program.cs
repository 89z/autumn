using C = System.Console;
using J = System.Text.Json.JsonSerializer;
using System.Linq;

class Program {
   static void Main() {
      string[] a = {"M", "a", "r", "c", "h"};
      // example 1
      var a1 = a.Skip(2).Take(2);
      // example 2
      var a2 = a.Skip(2);
      // print
      C.WriteLine(J.Serialize(a1) + J.Serialize(a2));
   }
}
