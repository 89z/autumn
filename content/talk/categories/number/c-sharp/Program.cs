using System;

class Program {
   static void Main() {
      // example 1
      double n1 = Math.Pow(7, 19);
      // example 2
      ulong n2 = 1;
      for (int n = 0; n < 19; n++) {
         n2 *= 7;
      }
      // print
      Console.WriteLine("{0} {1}", n1, n2);
   }
}
