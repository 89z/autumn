import asyncdispatch, httpclient, strformat
proc f1(n1, n2, n3: int64) {.async.} =
   echo &"total: {n1}, received: {n2}, speed: {n3 div 1000}k"
proc f2(s3: string) {.async.} =
   var o1 = newAsyncHttpClient()
   o1.onProgressChanged = f1
   await o1.downloadFile(s3, "100mb.file")
var s2 = "http://speedtest.lax.hivelocity.net/100mb.file"
waitFor f2(s2)
