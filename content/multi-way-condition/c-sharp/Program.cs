using System;

class Program {
   static void Main() {
      int n = 1 + 2;
      string s;

      switch (n) {
      case 5:
         s = "five";
         break;
      case 4:
      case 3:
         s = "four or three";
         break;
      default:
         s = n.ToString();
         break;
      }

      Console.WriteLine(s == "four or three");
   }
}
