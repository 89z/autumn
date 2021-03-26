import osproc
let a = ["google.com/search?tbm=vid&q=squarepusher"]
let p = startProcess("waterfox", args=a)
p.close
