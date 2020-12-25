using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;
using System.Linq;

class Program {
   static void Main() {
      string[] a = {"M", "a", "r", "c", "h"};
      // example 1
      var a1 = a.Skip(2).Take(2);
      // example 2
      var a2 = a.Skip(2);
      // print
      C.WriteLine(J.SerializeObject(a1) + J.SerializeObject(a2));
   }
}
