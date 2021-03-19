import strutils
let s = "\x0a\x0b\x0c\x0d"
let t = s.toHex
echo t == "0A0B0C0D"
