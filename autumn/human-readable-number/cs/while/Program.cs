using System;

class Program {
   static string NumberFormat(double n) {
      var (n2, n3) = (n, 0);
      while (n2 >= 1e3) {
         n2 /= 1e3;
         n3++;
      }
      return String.Format("{0:f3}", n2) + new[]{"", " k", " M", " G"}[n3];
   }

   static void Main() {
      var s = NumberFormat(9012345678);
      Console.WriteLine(s == "9.012 G");
   }
}
