import osproc
let o = "nim -v".execCmdEx
let s = o.output
stdout.write s
