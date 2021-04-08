import io = std.stdio;
import proc = std.process;

void main() {
   { // example 1
      auto c = proc.execute("dust");
      io.write(c.output);
   }
   { // example 2
      auto c = proc.execute(["dust", "-V"]);
      io.write(c.output);
   }
}
