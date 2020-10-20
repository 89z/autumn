using CommandLine;
using System;

class Program {
   class Options {
      [Option("start")]
      public string Start { get; set; }
      [Option("end")]
      public string End { get; set; }
      [Value(0, MetaName = "input", Required = true)]
      public string Input { get; set; }
   }

   static void Main(string[] a) {
      var s_start = "";
      var s_end = "";
      var s_input = "";
      Parser.Default.ParseArguments<Options>(a).WithParsed<Options>(o => {
         s_start = o.Start;
         s_end = o.End;
         s_input = o.Input;
      });
      Console.WriteLine(s_start + s_input + s_end);
   }
}
