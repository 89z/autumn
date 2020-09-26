using System;
using System.IO;

class Program {
   static void Main() {
      var s = "Program.cs";
      var s1 = Path.GetFullPath(s);
      Console.WriteLine(s1);
   }
}
