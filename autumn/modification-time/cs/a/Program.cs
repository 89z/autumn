using C = System.Console;
using F = System.IO.File;

class Program {
   static void Main() {
      var o = F.GetLastWriteTime("Program.cs");
      C.WriteLine(o);
   }
}
