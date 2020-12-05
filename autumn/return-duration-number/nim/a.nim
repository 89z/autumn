var a = initDateTime(25, mMar, 2015, 12, 0, 0, utc())
var b = initDateTime(1, mApr, 2017, 15, 0, 15, utc())
var ti = initTimeInterval(years = 2, weeks = 1, hours = 3, seconds = 15)
doAssert between(a, b) == ti
doAssert between(a, b) == -between(b, a)
