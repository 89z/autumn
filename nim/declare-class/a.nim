type Time = object
   hours: int

proc newTime(): Time =
   return Time(hours: 23)

proc duration(o: Time, minutes: int): int =
   return o.hours * 60 + minutes

let o = newTime()
let n = o.duration(59)
echo n == 1439
