import 'dart:io';

void main() {
   var o_uri = Uri.parse("https://speedtest.lax.hivelocity.net");
   var o_http = new HttpClient();
   o_http.getUrl(o_uri).then(
      (o_req) => o_req.close()
   ).then(
      (o_resp) => o_resp.pipe(stdout)
   );
   o_http.close();
}
