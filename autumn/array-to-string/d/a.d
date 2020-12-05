import array = std.array;
import io = std.stdio;

void main() {
   auto a = ["May", "June"];
   // example 1
   auto s1 = array.join(a, ",");
   // example 2
   auto s2 = array.join(a);
   // print
   io.writeln(s1 == "May,June" && s2 == "MayJune");
}
