using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "[\"sigma\", \"tau\"]";
      var a = J.Deserialize<string[]>(s);
      C.WriteLine(a);
   }
}
