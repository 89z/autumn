import io = std.stdio;
import url = urllibparse;

void main() {
   auto s = "month=March&day=Friday";
   auto m = url.parseQS(s);
   io.writeln(m);
}
