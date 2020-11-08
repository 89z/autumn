using System;

class Program {
   static void Main() {
      var o = new Uri("https://example.com/one?two=even");
      // example 1
      var s1 = o.DnsSafeHost;
      // example 2
      var s2 = o.LocalPath;
      // example 3
      var s3 = o.Query;
      // print
      Console.WriteLine(
         s1 == "example.com" && s2 == "/one" && s3 == "?two=even"
      );
   }
}
