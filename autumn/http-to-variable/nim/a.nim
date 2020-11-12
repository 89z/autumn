import httpclient
# -d:ssl
let in_s = "https://speedtest.lax.hivelocity.net"
let o = newHttpClient()
let out_s = o.getContent(in_s)
stdout.write out_s
