import string
# begin
def r64_encode(n_in):
   s_safe = string.printable[:62] + '-_'
   s_out = ''
   while n_in > 0:
      s_out += s_safe[n_in & 63]
      n_in >>= 6
   return s_out
# end
s1 = r64_encode(1234567890)
print(s1 == 'ibwB91')
