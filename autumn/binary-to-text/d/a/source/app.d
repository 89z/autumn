import botan.all: Pipe;
import botan.filters.b64_filt;
import std.stdio;

void main() {
   auto pipe = Pipe(new Base64Encoder);
   pipe.processMsg("\n\v\f\r");
   writeln(pipe.toString == "CgsMDQ==");
}
