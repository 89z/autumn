import conv = std.conv;
import io = std.stdio;

void main() {
   auto s = conv.hexString!"0A0B0C0D";
   io.writeln(s == "\n\v\f\r");
}
