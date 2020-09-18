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

n = 1577858399
o = Radix64()
# example 1
s1 = o.encode(n)
n1 = o.decode(s1)
# example 2
s2 = o.encode(n - 1)
n2 = o.decode(s2)
# print
print(s1 == '0T22KU', n1 == n, s2 == '0T22KT', n2 == n - 1)
