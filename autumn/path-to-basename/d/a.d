import
   io = std.stdio,
   path = std.path;

void main() {
   auto in_s = `C:\Windows\notepad.exe`;
   auto out_s = path.baseName(in_s);
   io.writeln(out_s == "notepad.exe");
}
