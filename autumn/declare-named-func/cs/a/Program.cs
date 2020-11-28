using System;

static class Program {
   // example 1
   static int add_f(this int n, int n1) {
      return n + n1;
   }
   // example 2
   static int sub_f(this int n, int n2) => n - n2;
   // print
   static void Main() {
      var n = 7;
      var n1 = n.add_f(1).add_f(1);
      var n2 = add_f(n,1).add_f(1);
      Console.WriteLine("{0} {1}", n1, n2);
   }
}
