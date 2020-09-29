using System;
using System.Web;

class Program {
   static void Main() {
      var m = HttpUtility.ParseQueryString("one=odd&two=even");
      foreach (var s_k in m.AllKeys) {
         var s_v = m[s_k];
         Console.WriteLine("{0} {1}", s_k, s_v);
      }
   }
}
