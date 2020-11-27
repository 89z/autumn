using System.IO;

class Test {
   static void Main() {
      var o = new System.DateTime(2000, 1, 1);
      File.SetLastWriteTime("Program.cs", o);
   }
}
