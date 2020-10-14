using System.CommandLine;
using System;

class Program {
   static void Main(string start, string cont, string s) {
      if (s == "") {
         Console.WriteLine(@"test [flags] <string>
--cont string
   contains
--start string
   starts with");
         Environment.Exit(1);
      }
      var b = s.StartsWith(start) && s.Contains(cont);
      Console.WriteLine(b);
   }
}
