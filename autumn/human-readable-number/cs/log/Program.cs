using System;

class Program {
   static void Main() {
      string[] suf = { "", " k", " M", " G" };
      long bytes = 1264;
      int place = Convert.ToInt32(Math.Floor(Math.Log(bytes, 1024)));
      double num = Math.Round(bytes / Math.Pow(1024, place), 3);
      var s = num.ToString() + suf[place];
      Console.WriteLine(s == "1.234 k");
   }
}
