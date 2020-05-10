import osproc
let o1 = startProcess("less", args=["-V"], options={poParentStreams})
discard o1.waitForExit
