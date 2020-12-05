import ar = std.array;
import io = std.stdio;
import re = std.regex;

void main() {
   auto o = re.matchAll("Sunday Monday", "..n");
   auto a = ar.array(o);
   io.writeln(a[0][0] == "Sun" && a[1][0] == "Mon");
}
