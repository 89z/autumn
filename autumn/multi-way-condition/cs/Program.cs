using System;

class Program {
   static void Main() {
      var (s, n) = ("\u263A", 0);

      switch (s) {
      case "A":
         n = 1;
         break;
      case "¡":
      case "¢":
         n = 2;
         break;
      default:
         n = s.Length;
         break;
      }

      Console.WriteLine(n);
   }
}
