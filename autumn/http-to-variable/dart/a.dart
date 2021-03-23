import 'dart:io';

void main() {
   var uri = Uri.parse("https://speedtest.lax.hivelocity.net");
   var http = new HttpClient();
   http.getUrl(uri).then(
      (req) => req.close()
   ).then(
      (res) => res.pipe(stdout)
   );
   http.close();
}
