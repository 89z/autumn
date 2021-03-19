import base64
let s = "CgsMDQ=="
let t = decode(s)
echo t == "\x0a\x0b\x0c\x0d"
