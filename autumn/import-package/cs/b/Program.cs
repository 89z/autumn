using C = System.Console;
using D = System.Collections.Generic.Dictionary<string, int>;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      D[] a = {
         new D{{"month", 12}, {"day", 31}},
         new D{{"month", 11}, {"day", 30}},
      };
      C.WriteLine(J.Serialize(a));
   }
}
