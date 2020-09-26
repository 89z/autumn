using System;
using System.IO;

class Program {
   static void Main() {
      var b = File.Exists("Program.cs");
      Console.WriteLine(b);
   }
}

