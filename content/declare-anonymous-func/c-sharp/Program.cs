using System;

class Program {
   // example 1
   static Func<int, int, bool> f1 = (n, n1) => {
      return n > n1;
   };
   // example 2
   static Func<int, int, bool> f2 = (n, n1) => n > n1;
   // print
   static void Main() {
      var b1 = f1(9, 8);
      var b2 = f2(9, 8);
      Console.WriteLine(b1 && b2);
   }
}
