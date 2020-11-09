using System.CommandLine.Invocation;
using System.CommandLine;
using System;

class Program {
   static void Main(string[] a) {
      var s_pre = "";
      var s_suf = "";
      var s_stem = "";

      var o = new RootCommand {
         new Option<string>("-p", description: "prefix"),
         new Option<string>("-s", description: "suffix"),
         new Argument<string>("stem")
      };

      void Execute(string p, string s, string stem) {
         s_pre = p;
         s_suf = s;
         s_stem = stem;
      }

      o.Handler = CommandHandler.Create<string, string, string>(Execute);
      o.Invoke(a);

      if (s_stem == "") {
         Environment.Exit(1);
      }

      Console.WriteLine(s_pre + s_stem + s_suf);
   }
}
