using System;

class Program {
   static void Main() {
      var s = "9.9";
      // example 1
      var n1 = float.Parse(s);
      // example 2
      var n2 = double.Parse(s);
      // print
      Console.WriteLine(n1 == 9.9f && n2 == 9.9);
   }
}
