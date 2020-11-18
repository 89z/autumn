using System;

class Program {
   static string NumberFormat(double n) {
      double n2 = n;
      int n3 = 0;
      while (n2 >= 1000) {
         n2 /= 1000;
         n3++;
      }
      return String.Format("{0:f3}", n2) + new[]{"", " k", " M", " G"}[n3];
   }

   static void Main() {
      var s = NumberFormat(9012345678);
      Console.WriteLine(s == "9.012 G");
   }
}
