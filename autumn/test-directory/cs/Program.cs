using C = System.Console;
using D = System.IO.Directory;

class Program {
   static void Main() {
      var b = D.Exists("obj");
      C.WriteLine(b);
   }
}

