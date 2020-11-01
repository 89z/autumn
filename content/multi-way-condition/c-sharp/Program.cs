using System;
class Program {
   static void Main() {
      int n = 46 / 10;
      string s;

      switch (n) {
      case 7:
         s = "seven";
         break;
      case 6:
      case 5:
         s = "six or five";
         break;
      default:
         s = n.ToString();
         break;
      }

      Console.WriteLine(s);
   }
}
