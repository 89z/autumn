import ae.sys.net, std.stdio;

void main() {
   auto s = "https://speedtest.lax.hivelocity.net";
   auto y = net.getFile(s);
   write(cast(char[])y);
}
