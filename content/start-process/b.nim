import osproc
let o = startProcess("less", args=["-V"], options={poParentStreams})
discard o.waitForExit
