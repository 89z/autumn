import httpclient
# -d:ssl
let s_in = "https://speedtest.lax.hivelocity.net"
let o = newHttpClient()
let s_out = o.getContent(s_in)
stdout.write s_out
