using System;

class Program {
   static void Main() {
      var n = 9.9;
      var s = n.ToString("#");
      Console.WriteLine(s == "10");
   }
}
