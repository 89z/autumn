using System;

class Program {
   static void Main() {
      var n1 = 1000;
      // example 1
      Console.WriteLine("{0:n}", n1);
      // example 2
      var s = $"{n1:n}";
      Console.WriteLine(s == "1,000.00");
   }
}
