import uri
let m = {"west": "left", "east": "right"}
let s = m.encodeQuery
echo s == "west=left&east=right"
