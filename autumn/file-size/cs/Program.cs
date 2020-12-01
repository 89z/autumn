using C = System.Console;
using IO = System.IO;

class Program {
   static void Main() {
      var o = new IO.FileInfo("Program.cs");
      var n = o.Length;
      C.WriteLine(n);
   }
}
