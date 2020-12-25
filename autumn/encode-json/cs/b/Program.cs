using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;

class Program {
   static void Main() {
      string[,] a = { { "one", "two" }, { "three", "four" } };
      string s = J.SerializeObject(a);
      C.WriteLine(s);
   }
}
