using System;

class Program {
   static void Main() {
      var s = "June";
      // example 1
      var s1 = s[2];
      // example 2
      var s2 = s.Substring(2, 1);
      // example 3
      var s3 = s.Substring(2);
      // print
      Console.WriteLine(s1 == 'n' && s2 == "n" && s3 == "ne");
   }
}
