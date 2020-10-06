using System;

class Program {
   static void Main() {
      // example 1
      var s = "May";
      var n1 = s.Length;
      // example 2
      var s2 = "📗";
      var n2 = s2.Length;
      // print
      Console.WriteLine("{0} {1}", n1 == 3, n2 == 2);
   }
}
