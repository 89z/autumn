import 'dart:io';

void main() async {
   var o_uri = Uri.parse("https://speedtest.lax.hivelocity.net");
   var o_http = new HttpClient();
   var o_req = await o_http.getUrl(o_uri);
   var o_resp = await o_req.close();
   o_resp.pipe(stdout);
   o_http.close();
}
