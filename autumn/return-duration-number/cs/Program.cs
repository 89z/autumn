using System;

class Program {
   static void Main() {
      var old_o = DateTime.Parse("2019-12-31");
      var new_o = DateTime.Now;
      var n = (new_o - old_o).TotalDays;
      Console.WriteLine(n);
   }
}
