using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      C.WriteLine(J.Serialize(a));
   }
}
