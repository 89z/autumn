def u_group(s_in):
   n_neg = ('-' + s_in).rfind('-')
   n_dot = (s_in + '.').find('.')
   s_out = ''
   for n_pos, s_pos in enumerate(s_in):
      if n_neg < n_pos and n_pos < n_dot and n_pos % 3 == n_dot % 3:
         s_out += ','
      s_out += s_pos
   return s_out
a1 = ['123', '1234', '-123', '-1234', '123.4', '1234.5']
for s1 in a1:
   print(u_group(s1))
