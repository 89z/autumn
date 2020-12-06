using System;

class Program {
   // example 1
   static bool f1(int n, int n1) {
      return n > n1;
   }
   // example 2
   static bool f2(int n, int n2) => n > n2;
   // print
   static void Main() {
      var b1 = f1(9, 8);
      var b2 = f2(9, 8);
      Console.WriteLine(b1 && b2);
   }
}
