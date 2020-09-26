using System;
using System.Collections;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      Array.Sort(a);
      foreach (var s in a) {
         Console.WriteLine(s);
      }
   }
}
