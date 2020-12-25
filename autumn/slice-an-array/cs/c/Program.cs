using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;

class Program {
   static void Main() {
      string[] a = {"M", "a", "r", "c", "h"};
      // example 1
      var a1 = a[2 .. 4];
      // example 2
      var a2 = a[2 ..];
      // print
      C.WriteLine(J.SerializeObject(a1) + J.SerializeObject(a2));
   }
}
