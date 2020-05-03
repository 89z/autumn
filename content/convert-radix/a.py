import string
s_dig = string.printable[:62] + '-_'
# example 1
def r64_encode(n_in):
   s_out = ''
   n_sh = n_in
   while n_sh > 0:
      n_ind = n_sh & 63
      s_out += s_dig[n_ind]
      n_sh >>= 6
   return s_out
s1 = r64_encode(1234567890)
# example 2
def r64_decode(s_in):
   n_ind = 0
   n_out = 0
   for n_sh in range(0, 36, 6):
      s_chr = s_in[n_ind]
      n_out |= s_dig.find(s_chr) << n_sh
      n_ind += 1
   return n_out
n1 = r64_decode('ibwB91')
# print
print(s1 == 'ibwB91', n1 == 1234567890)
