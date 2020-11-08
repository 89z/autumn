using System;
using System.IO;

class Program {
   static void Main() {
      var s = @"C:\Windows\notepad.exe";
      var s1 = Path.GetFileName(s);
      Console.WriteLine(s1 == "notepad.exe");
   }
}
