import httpclient
var s1 = "http://speedtest.lax.hivelocity.net"
var o1 = newHttpClient()
o1.downloadFile(s1, "index.html")
