using System;

class Program {
   static string NumberFormat(double n) {
      int n2 = 0;
      while (n > 1000) {
         n /= 1000;
         n2++;
      }
      return String.Format("{0:f3}", n) + new[]{"", " k", " M", " G"}[n2];
   }

   static void Main() {
      var s = NumberFormat(9012345678);
      Console.WriteLine(s == "9.012 G");
   }
}
