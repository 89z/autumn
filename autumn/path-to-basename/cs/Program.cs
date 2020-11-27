using C = System.Console;
using P = System.IO.Path;

class Program {
   static void Main() {
      var path_s = @"C:\Windows\notepad.exe";
      var file_s = P.GetFileName(path_s);
      C.WriteLine(file_s == "notepad.exe");
   }
}
