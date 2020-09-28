using System;
using System.Collections.Generic;

class Program {
   static void Main() {
      var a = new List<string>();
      a.Add("May");
      a.Add("June");
      foreach (var s in a) {
         Console.WriteLine(s);
      }
   }
}
