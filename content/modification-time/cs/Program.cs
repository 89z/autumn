using System;
using System.IO;

class Program {
   static void Main() {
      var o = File.GetLastWriteTime("Program.cs");
      Console.WriteLine(o);
   }
}
