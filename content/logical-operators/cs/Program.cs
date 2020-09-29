using System;

class Program {
   static void Main() {
      // example 1
      var b1 = ! false;
      // example 2
      var b2 = true && true;
      // example 3
      var b3 = false || true;
      // print
      Console.WriteLine("{0} {1} {2}", b1, b2, b3);
   }
}
