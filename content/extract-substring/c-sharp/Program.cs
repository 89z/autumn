using System;

class Program {
   static void Main() {
      var s = "May";
      // example 1
      var s1 = s[1];
      // example 2
      var s2 = s.Substring(1, 1);
      // print
      Console.WriteLine(s1 == 'a' && s2 == "a");
   }
}
