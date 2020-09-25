using System;

class Program {
   bool f(int n, int n1) => n > n1;
   static void Main() {
      var o = new Program();
      var b = o.f(9, 8);
      Console.WriteLine(b);
   }
}
