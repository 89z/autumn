using System;

class Program {
   static void Main() {
      var s = "May June";
      var a = s.Split();
      foreach (var s1 in a) {
         Console.WriteLine(s1);
      }
   }
}
