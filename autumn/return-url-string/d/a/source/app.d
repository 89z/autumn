import ae = ae.net.ietf.url;
import io = std.stdio;

void main() {
   auto s = "month=March&day=Friday";
   auto m = ae.decodeUrlParameters(s);
   io.writeln(m);
}
