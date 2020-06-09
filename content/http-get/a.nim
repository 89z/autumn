import httpclient
let s1 = "http://speedtest.lax.hivelocity.net"
let o1 = newHttpClient()
let s2 = o1.getContent(s1)
stdout.write s2
