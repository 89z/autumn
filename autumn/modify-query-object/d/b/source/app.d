import ae = ae.net.ietf.url;
import io = std.stdio;

void main() {
   auto m = ["month": "March", "day": "Friday"];
   auto s = ae.encodeUrlParameters(m);
   io.writeln(s);
}
