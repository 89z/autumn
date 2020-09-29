using System;

class Program {
   static void Main() {
      var s = "10";
      // example 1
      var n1 = int.Parse(s);
      // example 2
      var n2 = uint.Parse(s);
      // example 3
      var n3 = long.Parse(s);
      // example 4
      var n4 = ulong.Parse(s);
      // print
      Console.WriteLine(n1 == 10 && n2 == 10 && n3 == 10 && n4 == 10);
   }
}
