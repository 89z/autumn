using CommandLine;
using System;

class Program {
   class Options {
      [Option('c')]
      public string Contains { get; set; }
      [Option('p')]
      public string Prefix { get; set; }
   }
   static void Main(string[] a) {
      Parser.Default.ParseArguments<Options>(a).WithParsed<Options>(o => {
         Console.WriteLine("{0} {1}", o.Contains, o.Prefix);
      });
   }
}
