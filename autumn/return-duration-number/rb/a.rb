require 'date'
o = DateTime.parse('2020-05-04')
n = (DateTime.now - o).to_f
puts n
