using System;

class Program {
   static string NumberFormat(double n) {
      string[] a = { "", " k", " M", " G" };
      int n2 = 0;
      while (n >= 1024) {
         n /= 1024;
         n2++;
      }
      return String.Format("{0:f3}", n) + a[n2];
   }

   static void Main() {
      var s = NumberFormat(1264);
      Console.WriteLine(s == "1.234 k");
   }
}
