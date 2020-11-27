using Co = System.Console;
using Di = System.Diagnostics;
using Th = System.Threading.Thread;

class Program {
   static void Main() {
      var o = new Di.Stopwatch();
      o.Start();
      while (true) {
         Th.Sleep(10);
         Co.Write(
            "{0} m {1:00}.{2:00} s\r",
            o.Elapsed.Minutes,
            o.Elapsed.Seconds,
            o.Elapsed.Milliseconds / 10
         );
      }
   }
}
