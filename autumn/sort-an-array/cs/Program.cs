using A = System.Array;
using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      A.Sort(a);
      C.WriteLine(J.Serialize(a));
   }
}
