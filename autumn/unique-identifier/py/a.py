class Radix:
   sDigit = '0123456789abcdefghijklmnopqrstuvwxyz'

   def decode(self, sIn, nBase):
      return int(sIn, nBase)

   def encode(self, nIn, nBase):
      n = nIn // nBase
      s = self.sDigit[nIn % nBase]
      return self.encode(n, nBase) + s if n > 0 else s

o = Radix()
s = o.encode(1609480799, 36)
n = o.decode('qm8rbz', 36)
print(s == 'qm8rbz', n == 1609480799)
