using System;

class Program {
   static void Main() {
      int n = 10;
      string s;
      if (n > 0) {
         s = "+";
      } else if (n < 0) {
         s = "-";
      } else {
         s = "zero";
      }
      Console.WriteLine(s == "+");
   }
}
