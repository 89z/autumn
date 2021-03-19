import strutils
let s = "\n\v\f\r"
let t = s.toHex
echo t == "0A0B0C0D"
