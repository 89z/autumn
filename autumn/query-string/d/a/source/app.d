import io = std.stdio;
import web = vibe.inet.webform: FormFields;

void main() {
   FormFields m;
   web.parseURLEncodedForm("month=March&day=Friday", m);
   io.writeln(m);
}
