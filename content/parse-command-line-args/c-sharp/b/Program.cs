using System.CommandLine.Invocation;
using System.CommandLine;
using System;

class Program {
   static void Main(string[] a) {
      var s_start = "";
      var s_end = "";
      var s_input = "";

      var o = new RootCommand {
         new Option<string>("-start"),
         new Option<string>("-end"),
         new Argument<string>("input")
      };

      void Execute(string start, string end, string input) {
         s_start = start;
         s_end = end;
         s_input = input;
      }

      o.Handler = CommandHandler.Create<string, string, string>(Execute);
      o.Invoke(a);

      if (s_input == "") {
         Environment.Exit(2);
      }

      Console.WriteLine(s_start + s_input + s_end);
   }
}
