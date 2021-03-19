from base64 import nil
let s = "\n\v\f\r"
let t = base64.encode(s)
echo t == "CgsMDQ=="
