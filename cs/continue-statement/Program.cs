using System;

class Program {
   static void Main() {
      for (var n = 0; n < 10; n++) {
         if (n % 3 == 0) {
            continue;
         }
         Console.WriteLine(n);
      }
   }
}
