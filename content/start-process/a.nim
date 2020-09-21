import osproc

# example 1
echo "BEGIN"
discard execCmd("pipe")

# example 2
echo "BEGIN"
let o = startProcess("pipe", options={poParentStreams})
discard o.waitForExit
