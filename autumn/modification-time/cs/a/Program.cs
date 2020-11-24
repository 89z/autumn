using System.IO;
using System;

class Program {
   static void Main() {
      var o = File.GetLastWriteTime("Program.cs");
      Console.WriteLine(o);
   }
}
