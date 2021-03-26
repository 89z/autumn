import times

block: # example 1
   let t = initDateTime(31, mDec, 2020, 1, 2, 31, utc())
   echo t

block: # example 2
   let t = initDateTime(31, mDec, 2020, 1, 2, 31)
   echo t
