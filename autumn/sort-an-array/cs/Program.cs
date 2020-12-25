using A = System.Array;
using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      A.Sort(a);
      C.WriteLine(J.SerializeObject(a));
   }
}
