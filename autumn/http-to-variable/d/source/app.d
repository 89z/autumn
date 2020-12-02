import
   io = std.stdio,
   net = ae.sys.net;

void main() {
   auto url_s = "https://speedtest.lax.hivelocity.net";
   auto get_s = cast(string) net.getFile(url_s);
   io.write(get_s);
}
