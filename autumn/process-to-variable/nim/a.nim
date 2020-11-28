import osproc
let o = execCmdEx("nim -v")
let s = o.output
stdout.write s
