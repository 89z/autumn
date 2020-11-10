using System.IO;
using System;

class Program {
   static void Main() {
      var b = File.Exists("Program.cs");
      Console.WriteLine(b);
   }
}

