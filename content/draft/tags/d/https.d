import etc.c.curl, std.string;
void main() {
   auto s = "https://speedtest.lax.hivelocity.net";
   auto r = curl_easy_init();
   curl_easy_setopt(r, CurlOption.url, s.toStringz);
   curl_easy_perform(r);
}
