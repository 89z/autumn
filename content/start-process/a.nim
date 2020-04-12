import osproc
# example 1
discard execCmd("ag -V")
# example 2
let o1 = startProcess("ag", args=["-V"], options={poParentStreams})
discard o1.waitForExit
