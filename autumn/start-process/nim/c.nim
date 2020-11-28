import osproc
echo "BEGIN"
let o = startProcess("pipe", options={poParentStreams})
let n = o.waitForExit
echo n
