import uri
var u = initUri()
u.scheme = "http"
u.hostname = "nim-lang.org"
echo $u == "http://nim-lang.org"
