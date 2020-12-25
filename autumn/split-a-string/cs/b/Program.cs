using C = System.Console;
using J = Newtonsoft.Json.JsonConvert;

class Program {
   static void Main() {
      var s = "March";
      var a = s.ToCharArray();
      C.WriteLine(J.SerializeObject(a));
   }
}
