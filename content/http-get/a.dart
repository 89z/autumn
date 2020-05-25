main() async {
   var s = "http://speedtest.lax.hivelocity.net";
   var client = new HttpClient();
   var request = await client.getUrl(Uri.parse(s));
   var response = await request.close();
   response.pipe(stdout);
   client.close();
}
