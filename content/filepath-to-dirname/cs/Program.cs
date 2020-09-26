using System;
using System.IO;

class Program {
   static void Main() {
      var s = @"C:\Windows\notepad.exe";
      var s1 = Path.GetDirectoryName(s);
      Console.WriteLine(s1 == @"C:\Windows");
   }
}
