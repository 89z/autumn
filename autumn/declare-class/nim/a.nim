type Time = object
   hours: int

proc newTime(): Time =
   return Time(hours: 23)

proc duration(t: Time, minutes: int): int =
   return t.hours * 60 + minutes

let t = newTime()
let n = t.duration(59)
echo n == 1439
