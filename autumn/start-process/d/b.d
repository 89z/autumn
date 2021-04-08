import io = std.stdio;
import proc = std.process;

void main() {
   auto c = proc.executeShell("dust -V");
   io.writeln(c.output == "Dust 0.5.4\n");
}
