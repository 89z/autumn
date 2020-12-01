using Co = System.Console;
using Th = System.Threading.Thread;

class Program {
   static void Main() {
      Co.WriteLine("Sleep");
      Th.Sleep(500);
   }
}
