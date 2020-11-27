// dotnet add package System.CommandLine --prerelease
using C = System.Console;
using E = System.Environment;
using H = System.CommandLine.Invocation.CommandHandler;
using System.CommandLine;

class Program {
   static void Main(string[] a) {
      var pre_s = "";
      var suf_s = "";
      var stem_s = "";

      var o = new RootCommand {
         new Option<string>("-p", description: "prefix"),
         new Option<string>("-s", description: "suffix"),
         new Argument<string>("stem")
      };

      void Execute(string p, string s, string stem) {
         pre_s = p;
         suf_s = s;
         stem_s = stem;
      }

      o.Handler = H.Create<string, string, string>(Execute);
      o.Invoke(a);

      if (stem_s == "") {
         E.Exit(1);
      }

      C.WriteLine(pre_s + stem_s + suf_s);
   }
}
