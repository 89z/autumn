using System.Net;
using System;

class Program {
   static void Main() {
      var url_s = "https://speedtest.lax.hivelocity.net";
      var o = new WebClient();
      var res_s = o.DownloadString(url_s);
      Console.Write(res_s);
   }
}
