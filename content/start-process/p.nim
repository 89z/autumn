import osproc

# example 1
echo "begin"
discard execCmd("pipe")

# example 2
echo "begin"
let o = startProcess("pipe", options={poParentStreams})
discard o.waitForExit
