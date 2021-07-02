import botan.all: Pipe;
import botan.filters.hex_filt;
import std.stdio;

void main() {
   auto pipe = Pipe(new HexDecoder);
   pipe.processMsg("0A0B0C0D");
   writeln(pipe.toString == "\n\v\f\r");
}
