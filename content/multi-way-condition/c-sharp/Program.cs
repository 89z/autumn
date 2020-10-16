using System;

class Program {
   static void Main() {
      string s = "minute";
      int n;

      switch (s) {
      case "hour":
         n = 23;
         break;
      case "minute":
      case "second":
         n = 59;
         break;
      default:
         n = -1;
         break;
      }

      Console.WriteLine(n == 59);
   }
}
