import dlib.network.url: URL;
import std.stdio: writeln;

void main() {
   auto s = "https://example.com/one?two=even";
   auto o = URL!string(s);
   writeln(o.host == "example.com" && o.query == "two=even");
}
