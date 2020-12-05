import io = std.stdio;
import json = std.json: JSONOptions, JSONValue;

void main() {
   auto o = ["/", "📗"].JSONValue;
   // example 1
   auto s1 = json.toJSON(o);
   // example 2
   auto s2 = json.toJSON(o, true);
   // example 3
   auto s3 = json.toJSON(o, true, JSONOptions.doNotEscapeSlashes);
   // print
   io.writeln(s1, '\n', s2, '\n', s3);
}
