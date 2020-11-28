using System;

class Program {
   // example 1
   static int addf(int n, int n1) {
      return n + n1;
   }
   // example 2
   static int subf(int n, int n2) => n - n2;
   // print
   static void Main() {
      var n1 = addf(8, 1);
      var n2 = subf(8, 1);
      Console.WriteLine(n1 == 9 && n2 == 7);
   }
}
