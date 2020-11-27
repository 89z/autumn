using C = System.Console;
using P = System.IO.Path;

class Program {
   static void Main() {
      var s = P.GetExtension("Program.cs");
      C.WriteLine(s == ".cs");
   }
}
