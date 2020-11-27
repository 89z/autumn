using C = System.Console;
using D = System.Collections.Generic.Dictionary<string, int>;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      D[] a = {
         new D{{"year", 2019}, {"month", 12}},
         new D{{"year", 2018}, {"month", 11}},
      };
      var s = J.Serialize(a);
      C.WriteLine(s);
   }
}
