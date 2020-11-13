using System.Web;
using System;

class Program {
   static void Main() {
      var m = HttpUtility.ParseQueryString("one=odd&two=even");
      foreach (var s in m.AllKeys) {
         Console.WriteLine(s + ": " + m[s]);
      }
   }
}
