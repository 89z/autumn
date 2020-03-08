using Microsoft.CodeAnalysis.CSharp;
using System;
class Program {
   static void Main() {
      string s1 = SymbolDisplay.FormatLiteral("Sunday", quote: true);
      Console.WriteLine(s1);
   }
}
