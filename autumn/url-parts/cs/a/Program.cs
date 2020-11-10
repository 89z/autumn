using System;

class Program {
   static void Main() {
      var o = new Uri("https://example.com/one?two=even");
      Console.WriteLine(o.Host == "example.com");
   }
}
