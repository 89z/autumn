using C = System.Console;
using J = System.Text.Json.JsonSerializer;

class Program {
   static void Main() {
      // example 1
      int[] a1 = {10, 11};
      // example 2
      int[][] a2 = {
         new[]{10, 11},
         new[]{12, 13}
      };
      // print
      C.WriteLine(J.Serialize(a1) + J.Serialize(a2));
   }
}
