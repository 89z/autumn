import json

let s = """
{"month": 12, "day": 31}
"""

let m = s.parseJson
echo m["day"].getInt == 31
