class Radix:
   sDigit = '0123456789abcdefghijklmnopqrstuvwxyz'

   def decode(self, sIn, nBase):
      return int(sIn, nBase)

   def encode(self, nIn, nBase):
      sChar = self.sDigit[nIn % nBase]
      nIn //= nBase
      return self.encode(nIn, nBase) + sChar if nIn > 0 else sChar

o = Radix()
s = o.encode(1577858399, 36)
n = o.decode('q3ezbz', 36)
print(s == 'q3ezbz', n == 1577858399)
