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
      n.add_f(1).add_f(1);
      n.sub_f(1).sub_f(1);
      Console.WriteLine(n1 == 9 && n2 == 5);
   }
}
