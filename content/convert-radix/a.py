class Radix64:
   def __init__(self):
      self.s_dig = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz'

   def Encode(self, n_in):
      s_out = ''
      while True:
         s_out = self.s_dig[n_in % 64] + s_out
         n_in //= 64
         if n_in == 0:
            break
      return s_out

   def Decode(self, s_in):
      n_out = 0
      for s_chr in s_in:
         n_out = n_out * 64 + self.s_dig.find(s_chr)
      return n_out

import time
n = int(time.time())
o = Radix64()
s = o.Encode(n)
n2 = o.Decode(s)
print(n, s, n2 == n)
