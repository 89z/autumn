using System;
using System.IO;

class Program {
   static void Main() {
      var o = new StreamWriter("a.txt");
      o.WriteLine("May");
      o.Close();
   }
}
