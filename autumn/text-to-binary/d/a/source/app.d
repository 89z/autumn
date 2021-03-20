import botan.all: Pipe;
import botan.filters.b64_filt;
import std.stdio;

void main() {
   auto pipe = Pipe(new Base64Decoder);
   pipe.processMsg("CgsMDQ==");
   writeln(pipe.toString == "\n\v\f\r");
}
