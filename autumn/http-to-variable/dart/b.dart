import 'dart:io';

void main() async {
   var uri_o = Uri.parse("https://speedtest.lax.hivelocity.net");
   var http_o = new HttpClient();
   var req_o = await http_o.getUrl(uri_o);
   var resp_o = await req_o.close();
   resp_o.pipe(stdout);
   http_o.close();
}
