import etc.c.curl, std.string;
void main() {
   auto s1 = "http://speedtest.lax.hivelocity.net";
   auto r1 = curl_easy_init();
   curl_easy_setopt(r1, CurlOption.url, s1.toStringz);
   curl_easy_perform(r1);
}
