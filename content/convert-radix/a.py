class Radix64:
   def __init__(self):
      s = '-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz'
      self.s_dig = s

   def encode(self, n_in):
      s_out = ''
      while True:
         s_out = self.s_dig[n_in % 64] + s_out
         n_in //= 64
         if n_in == 0:
            break
      return s_out

   def decode(self, s_in):
      n_out = 0
      for s_chr in s_in:
         n_out = n_out * 64 + self.s_dig.find(s_chr)
      return n_out

import time
n = int(time.time())
o = Radix64()
# example 1
s1 = o.encode(n)
# example 2
n2 = o.decode(s1)
# print
print(n, s1, n2 == n)
