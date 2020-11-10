using System.IO;
using System;

class Program {
   static void Main() {
      var b = Directory.Exists("obj");
      Console.WriteLine(b);
   }
}

