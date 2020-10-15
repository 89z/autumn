using System.CommandLine.Invocation;
using System.CommandLine;
using System;

class Program {
   static void Main(string[] a) {
      int n_start = default;
      int n_len = default;
      string s_in = default;

      var o = new RootCommand {
         new Option<int>("-start"),
         new Option<int>("-len", getDefaultValue: () => 1),
         new Argument<string>("input")
      };

      void Execute(int start, int len, string input) {
         n_start = start;
         n_len = len;
         s_in = input;
      }

      o.Handler = CommandHandler.Create<int, int, string>(Execute);
      o.Invoke(a);
      var s_out = s_in.Substring(n_start, n_len);
      Console.WriteLine(s_out);
   }
}
