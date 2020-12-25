using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = { "one", "two" };
      string s = J.Serialize(a);
      C.WriteLine(s);
   }
}
