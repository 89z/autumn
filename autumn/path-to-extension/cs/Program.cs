using System.IO;
using System;

class Program {
   static void Main() {
      var s = "Program.cs";
      var s2 = Path.GetExtension(s);
      Console.WriteLine(s2 == ".cs");
   }
}
