using System;

class Program {
   static void Main() {
      var old_o = DateTime.Parse("2020-05-04");
      var new_o = DateTime.Now;
      var n = (new_o - old_o).TotalDays;
      Console.WriteLine(n);
   }
}
