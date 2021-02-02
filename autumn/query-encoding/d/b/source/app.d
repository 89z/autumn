import io = std.stdio;
import urllibparse;

void main() {
   auto m = ["month": "May", "day": "Friday"];
   auto s = urlEncode(m);
   io.writeln(s);
}
