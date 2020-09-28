using System;

class Program {
   static void Main() {
      // example 1
      var n1 = int.Parse("2147483647");
      // example 2
      var n2 = uint.Parse("4294967295");
      // example 3
      var n3 = long.Parse("9223372036854775807");
      // example 4
      var n4 = ulong.Parse("18446744073709551615");
      // print
      Console.WriteLine("{0} {1} {2} {3}", n1, n2, n3, n4);
   }
}
