import string
# begin
def r64_encode(n_div):
   s_safe = string.digits + string.ascii_letters + '-_'
   s_out = ''
   while n_div > 0:
      n_div, n_mod = divmod(n_div, 64)
      s_out += s_safe[n_mod]
   return s_out
# end
s1 = r64_encode(1234567890)
print(s1)
