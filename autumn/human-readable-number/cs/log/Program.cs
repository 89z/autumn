using System;

class Program {
   static string NumberFormat(int n) {
      string[] a = { "", " k", " M", " G" };
      double n2 = Math.Log(n, 1024);
      int n3 = Convert.ToInt32(Math.Floor(n2));
      return String.Format("{0:f3}", n / Math.Pow(1024, n3)) + a[n3];
   }

   static void Main() {
      var s = NumberFormat(1264);
      Console.WriteLine(s == "1.234 k");
   }
}
