import
   io = std.stdio,
   dlib.network.url: URL;

void main() {
   auto s = "https://example.com/one?two=even";
   auto o = URL!string(s);
   io.writeln(o.host == "example.com" && o.query == "two=even");
}
