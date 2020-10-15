using System;

class Program {
   static void Main() {
      var s = "11";
      // example 1
      var n1 = int.Parse(s);
      // example 2
      var n2 = uint.Parse(s);
      // example 3
      var n3 = long.Parse(s);
      // example 4
      var n4 = ulong.Parse(s);
      // print
      Console.WriteLine(n1 == 11 && n2 == 11 && n3 == 11 && n4 == 11);
   }
}
