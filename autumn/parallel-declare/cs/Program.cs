using System;

class Program {
   static void Main() {
      // example 1
      string s1, s1a;
      s1 = "May";
      s1a = "June";
      // example 2
      var (s2, s2a) = ("May", "June");
      // example 3
      string s3 = "May", s3a = "June";
      // print
      Console.WriteLine(s1 + s1a + s2 + s2a + s3 + s3a);
   }
}
