import osproc
let a = ["google.com/search?tbm=vid&q=squarepusher"]
let o = startProcess("waterfox", args=a)
o.close
