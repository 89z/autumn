import std.json, std.stdio;

void main() {
   auto v = ["/", "📗"].JSONValue;
   { // example 1
      auto s = v.toJSON;
      s.writeln;
   }
   { // example 2
      auto s = v.toJSON(true);
      s.writeln;
   }
   { // example 3
      auto s = v.toJSON(true, JSONOptions.doNotEscapeSlashes);
      s.writeln;
   }
}
