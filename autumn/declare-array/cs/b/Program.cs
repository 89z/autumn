using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      // example 1
      var a1 = new[]{10, 11};
      // example 2
      var a2 = new[]{
         new[]{10, 11},
         new[]{12, 13}
      };
      // print
      C.WriteLine(J.Serialize(a1) + J.Serialize(a2));
   }
}
