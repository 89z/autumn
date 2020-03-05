library('RCurl')
webpage <- getURL("http://speedtest.lax.hivelocity.net")
cat(webpage)
