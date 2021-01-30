import io = std.stdio;
import web = vibe.inet.webform;

void main() {
   auto m = ["month": "March", "day": "Friday"];
   auto s = web.urlEncode(m);
   io.writeln(s);
}
