import
   io = std.stdio,
   json = std.json: JSONValue;

void main() {
   auto o = ["/", "📗"].JSONValue;
   auto s = json.toJSON(o);
   io.writeln(s);
}
