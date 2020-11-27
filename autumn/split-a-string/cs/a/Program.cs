using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "M a r c h";
      // example 1
      var a1 = s.Split(" ");
      // example 2
      var a2 = s.Split();
      // print
      C.WriteLine(J.Serialize(a1) + J.Serialize(a2));
   }
}
