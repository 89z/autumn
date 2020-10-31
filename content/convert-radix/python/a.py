class Radix36:
   digit = '0123456789abcdefghijklmnopqrstuvwxyz'

   def encode(self, n):
      if n == 0: return ''
      return self.encode(n // 36) + self.digit[n % 36]

   def decode(self, s):
      return int(s, 36)

o = Radix36()
s = o.encode(1577858399)
n = o.decode('q3ezbz')
print(s == 'q3ezbz', n == 1577858399)
