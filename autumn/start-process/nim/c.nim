import osproc
let p = startProcess(r"C:\Windows\notepad", options={poParentStreams})
let n = p.waitForExit
echo n
