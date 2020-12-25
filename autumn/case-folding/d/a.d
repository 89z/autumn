import io = std.stdio;
import uni = std.uni;

void main() {
   auto s = "March";
   auto t = uni.toUpper(s);
   io.writeln(t == "MARCH");
}
