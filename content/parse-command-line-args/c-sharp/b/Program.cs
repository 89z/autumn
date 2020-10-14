using System.CommandLine.Invocation;
using System.CommandLine;
using System;

class Program {
   static void Main(string[] a) {
      var rootCommand = new RootCommand {
         new Option<int>("-year")
      };
      var n_year = 0;
      rootCommand.Handler = CommandHandler.Create<int>((year) => {
         n_year = year;
      });
      rootCommand.Invoke(a);
      Console.WriteLine(n_year);
   }
}
