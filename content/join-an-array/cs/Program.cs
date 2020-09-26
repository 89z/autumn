using System;

class Program {
   static void Main() {
      string[] a = {"May", "June"};
      var s = String.Join(",", a);
      Console.WriteLine(s == "May,June");
   }
}
