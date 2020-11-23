import
   ae.sys.net,
   std.stdio;

void main() {
   auto url_s = "http://speedtest.lax.hivelocity.net";
   auto get_s = cast(string) net.getFile(url_s);
   get_s.write;
}
