type date = ref object
   month, day: int

proc newDate(): date =
   return date(month: 12, day: 31)

proc year(d: date) =
   d.month = 1
   d.day = 1

let d = newDate()
d.year
echo d.day == 1
