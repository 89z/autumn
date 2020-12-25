using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "May,June,July";
      var a = s.Split(",");
      C.WriteLine(J.Serialize(a));
   }
}
