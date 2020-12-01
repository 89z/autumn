using C = System.Console;

class Program {
   static void Main() {
      C.WriteLine("example 1");
      var n1 = 10;
      while (n1 < 20) {
         C.WriteLine(n1);
         n1++;
      }

      C.WriteLine("example 2");
      var n2 = 10;
      while (true) {
         if (n2 > 19) {
            break;
         }
         C.WriteLine(n2);
         n2++;
      }
   }
}
