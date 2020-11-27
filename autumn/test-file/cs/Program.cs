using C = System.Console;
using F = System.IO.File;

class Program {
   static void Main() {
      var b = F.Exists("Program.cs");
      C.WriteLine(b);
   }
}

