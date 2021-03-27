type date = ref object
   month, day: int

proc newDate(): date =
   return date(month: 1, day: 1)

proc addDay(d: date) =
   d.day = d.day + 1

let d = newDate()
d.addDay
echo d.day == 2
