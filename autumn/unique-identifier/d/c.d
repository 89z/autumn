import id = std.uuid;
import io = std.stdio;

void main() {
   auto s = id.randomUUID();
   io.writeln(s);
}
