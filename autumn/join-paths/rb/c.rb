require 'pathname'
pn = Pathname.new('south') / 'north'
s = pn.to_s
p s == 'south/north'
