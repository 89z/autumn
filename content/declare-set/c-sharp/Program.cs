using System.Collections.Generic;
using System;
using static System.Text.Json.JsonSerializer;
 
class Program {
   static void Main() {
      // example 1
      var t1 = new HashSet<string>{};
      // example 2
      var t2 = new HashSet<string>{"May", "June"};
      // print
      Console.WriteLine("{0} {1}", Serialize(t1), Serialize(t2));
   }
}
