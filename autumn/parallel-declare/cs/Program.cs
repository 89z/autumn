using System;

class Program {
   static void Main() {
      // example 1
      string s1, s1a;
      s1 = "March";
      s1a = "April";
      // example 2
      var (s2, n2) = ("year", 2019);
      // print
      Console.WriteLine("{0} {1} {2} {3}", s1, s1a, s2, n2);
   }
}
