import json

let s = """
{"month": 12, "day": 31}
"""

let m = s.parseJson
let n = m["day"].getInt
echo n == 31
