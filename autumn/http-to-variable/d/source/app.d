import io = std.stdio;
import net = ae.sys.net;

void main() {
   auto url = "https://speedtest.lax.hivelocity.net";
   auto get = cast(string) net.getFile(url);
   io.write(get);
}
