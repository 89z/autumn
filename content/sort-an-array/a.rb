a = %w(b A a B)

a.sort
a.sort!
a.sort do |x, y|
  x.downcase <=> y.downcase
end
