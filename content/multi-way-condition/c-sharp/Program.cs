using System;

class Program {
   static void Main() {
      int n = 10;
      string s;
      switch (n) {
      case 12:
         s = "all";
         break;
      case 11:
      case 10:
         s = "some";
         break;
      default:
         s = "none";
         break;
      }
      Console.WriteLine(s == "some");
   }
}
