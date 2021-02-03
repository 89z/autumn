import io = std.stdio;
import urllibparse;

void main() {
   URLSplitResult u;
   u.netloc = "dpldocs.info";
   io.writeln(u);
}
