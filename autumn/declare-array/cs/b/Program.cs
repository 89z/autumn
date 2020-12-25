using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;

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
      C.WriteLine(J.SerializeObject(a1) + J.SerializeObject(a2));
   }
}
