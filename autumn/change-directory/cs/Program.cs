using C = System.Console;
using D = System.IO.Directory;

class Program {
   static void Main() {
      D.SetCurrentDirectory("..");
      var s = D.GetCurrentDirectory();
      C.WriteLine(s);
   }
}
