using System;

class Program {
   static void Main() {
      char c = '\x43';
      int n;

      switch (c) {
      case 'A':
         n = 0x41;
         break;
      case 'B':
      case 'b':
         n = 0x42;
         break;
      default:
         n = (int) c;
         break;
      }

      Console.WriteLine(n == 0x43);
   }
}
