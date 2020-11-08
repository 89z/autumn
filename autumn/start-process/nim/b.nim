import osproc
# example 1
let a = ["google.com/search?tbm=vid&q=squarepusher"]
discard startProcess("waterfox", args=a)
# example 2
echo "BEGIN"
let o = startProcess("pipe", options={poParentStreams})
discard o.waitForExit
