import json

let src = """
{"U2": {"Boy": ["Twilight", "I Will Follow"]}}
"""

let m = src.parseJson
let dst = m["U2"]["Boy"][0].getStr
echo dst == "Twilight"
