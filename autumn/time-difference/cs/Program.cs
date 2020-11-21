using System.Diagnostics;
using System.Threading;
using System;

class Program {
   static void Main() {
      var o = new Stopwatch();
      o.Start();
      Thread.Sleep(1234);
      Console.WriteLine("{0} {1}", o.Elapsed.Seconds, o.Elapsed.Milliseconds);
      Thread.Sleep(1234);
      Console.WriteLine("{0} {1}", o.Elapsed.Seconds, o.Elapsed.Milliseconds);
   }
}
