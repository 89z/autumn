import io = std.stdio;
import urllibparse;

void main() {
   URLSplitResult u = {
      netloc: "netloc.info", path: "/path", query: "month=May"
   };
   auto s = urlUnsplit(u);
   io.writeln(s);
}
