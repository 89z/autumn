using System.IO;
using System;

class Program {
   static void Main() {
      Directory.SetCurrentDirectory("..");
      var s = Directory.GetCurrentDirectory();
      Console.WriteLine(s);
   }
}
