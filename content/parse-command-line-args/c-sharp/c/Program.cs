using System.CommandLine.DragonFruit;
using System;

class Program {
   static void Main(string argument, int start, int len = 1) {
      var s_out = argument.Substring(start, len);
      Console.WriteLine(s_out);
   }
}
