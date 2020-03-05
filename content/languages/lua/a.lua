r1 = require('curl')
s1 = 'http://speedtest.lax.hivelocity.net/100mb.file'
r2 = io.open('100mb.file', 'w')
r1.easy{noprogress = false, url = s1, writefunction = r2}:perform()
