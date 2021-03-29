import io = std.stdio;
import urllibparse;

void main() {
   auto s = "http://netloc.info/path?west=left";
   auto a = urlSplit(s);
   io.writeln(a);
}
