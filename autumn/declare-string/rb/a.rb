# slash
s = <<eof
south\\north
eof
puts s
s = <<'eof'
south\north
eof
puts s

# quote
s = <<eof
south"north
south'north
eof
puts s

# newline
s = <<eof
south
north
eof
puts s
