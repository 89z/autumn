using System;

class Program {
   // example 1
   Func<int, int, bool> f1 = (n, n1) => {
      return n > n1;
   };
   // example 2
   Func<int, int, bool> f2 = (n, n2) => n > n2;
   // print
   static void Main() {
      var o = new Program();
      var b1 = o.f1(9, 8);
      var b2 = o.f2(9, 8);
      Console.WriteLine(b1 && b2);
   }
}
