using System;

class Program {
   static void Main() {
      var s = "March";
      var s2 = s[2 .. 3];
      Console.WriteLine(s2 == "r");
   }
}
