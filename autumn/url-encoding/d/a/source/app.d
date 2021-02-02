import io = std.stdio;
import urllibparse;

void main() {
   auto s = "http://netloc.info/path?month=May";
   auto a = urlSplit(s);
   io.writeln(a);
}
