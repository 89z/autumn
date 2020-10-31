class Radix36:
   digit = '0123456789abcdefghijklmnopqrstuvwxyz'

   def decode(self, s):
      return int(s, 36)

   def encode(self, n):
      result = []
      while num != 0:
         num, d = divmod(num, 36)
         result += digits[d]
      return ''.join(result[::-1])
