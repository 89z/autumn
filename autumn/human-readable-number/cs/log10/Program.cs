using System;

class Program {
   static string NumberFormat(double n) {
      int n2 = Convert.ToInt32(Math.Log10(n)) / 3;
      n /= Math.Pow(1000, n2);
      return String.Format("{0:f3}", n) + new[]{"", " k", " M", " G"}[n2];
   }

   static void Main() {
      var s = NumberFormat(1234);
      Console.WriteLine(s == "1.234 k");
   }
}
