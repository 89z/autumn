type Radix = object
   sDigit: string

proc newRadix(): Radix =
   return Radix(sDigit: "0123456789abcdefghijklmnopqrstuvwxyz")

proc encode(o: Radix, nIn, nBase: int): string =
   let n = nIn div nBase
   let c = o.sDigit[nIn mod nBase]
   return if n > 0: o.encode(n, nBase) & c else: $(c)

proc decode(o: Radix, sIn: string, nBase: int): int =
   let s = sIn[0 ..< ^1]
   let n = o.sDigit.find(sIn[^1])
   return if s != "": o.decode(s, nBase) * nBase + n else: n

let o = newRadix()
let s = o.encode(1609480799, 36)
let n = o.decode("qm8rbz", 36)
echo s == "qm8rbz" and n == 1609480799
