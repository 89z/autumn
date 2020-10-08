using System.Collections.Generic;
using System;

class Program {
   static void Main() {
      // example 1
      var m1 = new Dictionary<string, int>();
      // example 2
      var m2 = new Dictionary<string, int>{};
      // print
      Console.WriteLine("{0}\n{1}", m1, m2);
   }
}
