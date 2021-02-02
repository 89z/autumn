import io = std.stdio;
import urllibparse;

void main() {
   auto s = "month=March&day=Friday";
   auto m = parseQS(s);
   io.writeln(m);
}
