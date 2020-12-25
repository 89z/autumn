# example 1
s, t = "\u00B5", "\u03BC" # MICRO SIGN, GREEK SMALL LETTER MU
s, t = s.downcase(:fold), t.downcase(:fold)
puts s == t

# example 2
s, t = "\u006B", "\u212A" # LATIN SMALL LETTER K, KELVIN SIGN
s, t = s.downcase(:fold), t.downcase(:fold)
puts s == t
