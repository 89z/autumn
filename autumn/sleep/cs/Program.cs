using Co = System.Console;
using Th = System.Threading.Thread;

class Program {
   static void Main() {
      while (true) {
         Co.WriteLine("Sleep");
         Th.Sleep(500);
      }
   }
}
