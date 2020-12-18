import array = std.array;
import io = std.stdio;

void main() {
   auto s = "May,June";
   auto a = array.split(s, ",");
   io.writeln(a);
}
