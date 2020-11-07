using System;

static class Program {
   // example 1
   static bool f1(this int n, int n1) {
      return n > n1;
   }
   // example 2
   static bool f2(this int n, int n2) => n > n2;
   // print
   static void Main() {
      var n = 9;
      var b1 = n.f1(8);
      var b2 = n.f2(8);
      Console.WriteLine(b1 && b2);
   }
}
