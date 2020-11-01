using System;

class Program {
   static void Main() {
      int n = 15 / 10;
      string s;

      switch (n) {
      case 3:
         s = "three";
         break;
      case 2:
      case 1:
         s = "two or one";
         break;
      default:
         s = n.ToString();
         break;
      }

      Console.WriteLine(s == "two or one");
   }
}
