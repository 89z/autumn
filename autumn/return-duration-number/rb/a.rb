require 'date'
o = DateTime.parse('2020-12-31')
n = (DateTime.now - o).to_f
puts n
