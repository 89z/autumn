import
   io = std.stdio,
   net = dlib.network.url;

void main() {
   auto s = "https://example.com/one?two=even";
   auto o = net.URL!string(s);
   io.writeln(o.host == "example.com" && o.query == "two=even");
}
