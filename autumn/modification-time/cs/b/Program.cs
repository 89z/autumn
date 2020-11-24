using System.IO;
using System;

class Test {
   public static void Main() {
      var o = new DateTime(2000, 1, 1);
      File.SetLastWriteTime("Program.cs", o);
   }
}
