using System;
using System.Collections.Generic;
 
class Program {
   static void Main() {
      var t = new HashSet<string>{"May", "June"};
      foreach (var s in t) {
         Console.WriteLine(s);
      }
   }
}
