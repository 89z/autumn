using System;

class Program {
   static void Main() {
      // example 1
      string s, s1;
      s = "month";
      s1 = "month";
      // example 2
      var (s2, n) = ("month", 12);
      // print
      Console.WriteLine("{0} {1} {2} {3}", s, s1, s2, n);
   }
}
