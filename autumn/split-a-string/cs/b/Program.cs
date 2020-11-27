using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "March";
      var a = s.ToCharArray();
      C.WriteLine(J.Serialize(a));
   }
}
