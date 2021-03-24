require 'pathname'
pn = Pathname.new('south').join('north')
s = pn.to_s
p s == 'south/north'
