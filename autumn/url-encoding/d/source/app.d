import ae.net.http.common: HttpRequest;
import io = std.stdio;

void main() {
   auto u = new HttpRequest;
   u.host = "github.com";
   auto s = u.url();
   io.writeln(s);
}
