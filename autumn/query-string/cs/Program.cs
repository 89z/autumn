using C = System.Console;
using H = System.Web.HttpUtility;

class Program {
   static void Main() {
      var m = H.ParseQueryString("one=odd&two=even");
      foreach (var s in m.AllKeys) {
         C.WriteLine(s + ": " + m[s]);
      }
   }
}
