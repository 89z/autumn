import io = std.stdio;
import urllibparse;

void main() {
   URLSplitResult u;
   u.query = "month=May&day=Friday";
   io.writeln(u);
}
