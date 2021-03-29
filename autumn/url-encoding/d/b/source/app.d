import io = std.stdio;
import urllibparse;

void main() {
   URLSplitResult u = {
      netloc: "netloc.info", path: "/path", query: "west=left"
   };
   auto s = urlUnsplit(u);
   io.writeln(s);
}
