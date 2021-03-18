import dig = std.digest;
import io = std.stdio;

void main() {
   ubyte[] a = [10, 11, 12, 13];
   string s = dig.toHexString(a);
   io.writeln(s == "0A0B0C0D");
}
