import io = std.stdio;
import url = urllibparse;

void main() {
   auto m = ["month": "May", "day": "Friday"];
   auto s = url.urlEncode(m);
   io.writeln(s);
}
