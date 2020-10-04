using System;

class Program {
   // example 1
   bool f1(int n, int n1) {
      return n > n1;
   }
   // example 2
   bool f2(int n, int n1) => n > n1;
   // print
   static void Main() {
      var o = new Program();
      var b1 = o.f1(9, 8);
      var b2 = o.f2(9, 8);
      Console.WriteLine(b1 && b2);
   }
}
