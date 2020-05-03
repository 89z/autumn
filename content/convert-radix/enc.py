import string
# begin
def r64_encode(n_div):
   s_safe = string.printable[:62] + '-_'
   s_out = ''
   while n_div > 0:
      s_out += s_safe[n_div & 63]
      n_div >>= 6
   return s_out
# end
s1 = r64_encode(1234567890)
print(s1 == 'ibwB91')
