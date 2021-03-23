import 'dart:io';

void main() async {
   var uri = Uri.parse("https://speedtest.lax.hivelocity.net");
   var http = new HttpClient();
   var req = await http.getUrl(uri);
   var res = await req.close();
   res.pipe(stdout);
   http.close();
}
