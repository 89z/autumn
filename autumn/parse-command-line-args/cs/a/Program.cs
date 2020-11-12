using CommandLine;
using System;

class Program {
   class Options {
      [Option('p', HelpText = "prefix")]
      public string Prefix { get; set; }
      [Option('s', HelpText = "suffix")]
      public string Suffix { get; set; }
      [Value(0, MetaName = "stem", Required = true)]
      public string Stem { get; set; }
   }

   static void Main(string[] a) {
      var pre_s = "";
      var suf_s = "";
      var stem_s = "";
      Parser.Default.ParseArguments<Options>(a).WithParsed<Options>(o => {
         pre_s = o.Prefix;
         suf_s = o.Suffix;
         stem_s = o.Stem;
      });
      Console.WriteLine(pre_s + stem_s + suf_s);
   }
}
