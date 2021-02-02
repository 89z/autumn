import uri
let m = {"month": "May", "day": "Friday"}
let s = encodeQuery(m)
echo s == "month=May&day=Friday"
