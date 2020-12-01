import json

let in_s = """
{"U2": {"Boy": ["Twilight", "I Will Follow"]}}
"""

let m = in_s.parseJson
let out_s = m["U2"]["Boy"][0].getStr
echo out_s == "Twilight"
