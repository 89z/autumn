import strutils
let s = "0A0B0C0D"
let t = s.parseHexStr
echo t == "\x0a\x0b\x0c\x0d"
