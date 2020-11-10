import dlib.network.url, std.stdio;

void main() {
   auto o = URL!string("https://example.com/one?two=even");
   writeln(o.host == "example.com" && o.query == "two=even");
}
