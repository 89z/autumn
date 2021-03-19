import strutils
let s = "0A0B0C0D"
let t = s.parseHexStr
echo t == "\n\v\f\r"
