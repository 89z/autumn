require 'csv'

def f(s_in)
   a_head = []
   a_out = []
   CSV.new(s_in).each.with_index do |a_row, n_row|
      if n_row == 0
         a_head = a_row
         next
      end
      a_out << a_head.zip(a_row).to_h
   end
   return a_out
end

s = 'Month,Day
January,Sunday
February,Monday'

a = f(s)
p a
