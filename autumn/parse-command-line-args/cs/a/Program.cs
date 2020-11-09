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
      var s_pre = "";
      var s_suf = "";
      var s_stem = "";
      Parser.Default.ParseArguments<Options>(a).WithParsed<Options>(o => {
         s_pre = o.Prefix;
         s_suf = o.Suffix;
         s_stem = o.Stem;
      });
      Console.WriteLine(s_pre + s_stem + s_suf);
   }
}
