using C = System.Console;
using P = System.IO.Path;

class Program {
   static void Main() {
      var path_s = @"C:\Windows\notepad.exe";
      var dir_s = P.GetDirectoryName(path_s);
      C.WriteLine(dir_s == @"C:\Windows");
   }
}
