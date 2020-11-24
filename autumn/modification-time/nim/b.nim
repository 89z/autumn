import
   os,
   times

let o = fromUnix(400_000_000)
setLastModificationTime("a.nim", o)
