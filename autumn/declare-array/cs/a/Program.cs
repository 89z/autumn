using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;

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
      C.WriteLine(J.SerializeObject(a1) + J.SerializeObject(a2));
   }
}
