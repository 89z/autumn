using C = System.Console;
using N = System.Net;

class Program {
   static void Main() {
      var o = new N.WebClient();
      var url_s = "http://speedtest.lax.hivelocity.net";
      var res_s = o.DownloadString(url_s);
      C.Write(res_s);
   }
}
