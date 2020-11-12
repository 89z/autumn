import 'dart:io';

void main() {
   var uri_o = Uri.parse("https://speedtest.lax.hivelocity.net");
   var http_o = new HttpClient();
   http_o.getUrl(uri_o).then(
      (req_o) => req_o.close()
   ).then(
      (resp_o) => resp_o.pipe(stdout)
   );
   http_o.close();
}
