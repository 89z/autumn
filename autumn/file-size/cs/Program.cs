using C = System.Console;
using G = System.Collections.Generic;
using J = System.Text.Json.JsonSerializer;
using System.IO

class Program {
   static void Main() {
      var filePath = "Sample.txt";
      var fileInfo = new FileInfo(filePath);
      var _x = fileInfo.Length;
   }
}
