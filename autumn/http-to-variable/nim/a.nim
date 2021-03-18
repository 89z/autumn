import httpclient
# -d:ssl
let s = "https://speedtest.lax.hivelocity.net"
let t = newHttpClient().getContent(s)
stdout.write t
