using System;

class Program {
   static void Main() {
      string[] sizes = { " ", "K", "M", "G" };
      double len = 1264;
      int order = 0;
      while (len >= 1024 && order < sizes.Length - 1) {
         order++;
         len = len/1024;
      }
      string result = String.Format("{0:0.###} {1}", len, sizes[order]);
      Console.WriteLine(result == "1.234 K");
   }
}
