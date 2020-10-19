using CommandLine;
using System;

class Program {
   class Options {
      [Option('a', HelpText = "start")]
      public int Start { get; set; }
      [Option('b', Default = 1, HelpText = "length")]
      public int Length { get; set; }
      [Value(0, MetaName = "input", Required = true)]
      public string Input { get; set; }
   }

   static void Main(string[] a) {
      var n_start = 0;
      var n_len = 0;
      var s_in = "";

      Parser.Default.ParseArguments<Options>(a).WithParsed<Options>(o => {
         n_start = o.Start;
         n_len = o.Length;
         s_in = o.Input;
      });

      var s_out = s_in.Substring(n_start, n_len);
      Console.WriteLine(s_out);
   }
}
