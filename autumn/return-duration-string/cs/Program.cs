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
            "{0} s {1:000} ms\r", o.Elapsed.Seconds, o.Elapsed.Milliseconds
         );
      }
   }
}
