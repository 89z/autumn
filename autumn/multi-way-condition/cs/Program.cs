using System;

class Program {
   static void Main() {
      int n = 0x63 - 0x20;
      char c;

      switch (n) {
      case 0x41:
         c = 'A';
         break;
      case 0x42:
      case 0x62:
         c = 'B';
         break;
      default:
         c = (char) n;
         break;
      }

      Console.WriteLine(c == 'C');
   }
}
