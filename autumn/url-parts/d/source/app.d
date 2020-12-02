import dlib.network.url: URL;
import io = std.stdio;

void main() {
   auto s = "https://example.com/one?two=even";
   auto o = URL!string(s);
   io.writeln(o.host == "example.com" && o.query == "two=even");
}
