using System.Net;
using System;

class Program {
   static void Main() {
      var s = "https://speedtest.lax.hivelocity.net";
      var o = new WebClient();
      var s1 = o.DownloadString(s);
      Console.Write(s1);
   }
}
