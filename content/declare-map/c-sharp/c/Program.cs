using System;
using System.Collections.Generic;

namespace Autumn {
   using Dict = Dictionary<string, int>;
   class Program {
      static void Main() {
         var m = new Dict {
            {"year", 2019}, {"month", 12}
         };
         Console.WriteLine(m);
      }
   }
}
