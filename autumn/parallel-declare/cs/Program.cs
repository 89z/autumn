using System;

class Program {
   static void Main() {
      // example 1
      string s, s1;
      s = "March";
      s1 = "April";
      // example 2
      var (s2, n) = ("year", 2019);
      // print
      Console.WriteLine("{0} {1} {2} {3}", s, s1, s2, n);
   }
}
