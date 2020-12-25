using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;

class Program {
   static void Main() {
      var s = "May,June,July";
      var a = s.Split(",");
      C.WriteLine(J.SerializeObject(a));
   }
}
