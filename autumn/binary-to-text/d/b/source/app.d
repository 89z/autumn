import botan.all: Pipe;
import botan.filters.hex_filt;
import std.stdio;

void main() {
   auto pipe = Pipe(new HexEncoder);
   pipe.processMsg("\n\v\f\r");
   writeln(pipe.toString == "0A0B0C0D");
}
