import io = std.stdio;
import str = std.string;

void main() {
   auto s = "March";
   auto t = str.toUpper(s);
   io.writeln(t == "MARCH");
}
