using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      var s = "[10, 11]";
      var a = J.Deserialize<int[]>(s);
      C.WriteLine(a[0] == 10);
   }
}
