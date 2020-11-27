// dotnet add package System.CommandLine.DragonFruit --prerelease
using C = System.Console;
using System.CommandLine.DragonFruit;

class Program {
   static void Main(string argument, string prefix, string suffix) {
      C.WriteLine(prefix + argument + suffix);
   }
}
