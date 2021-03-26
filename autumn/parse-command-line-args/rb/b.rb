require 'optparse'
m = {p: '', s: ''}
o = OptionParser.new
o.on('-p string', 'prefix')
o.on('-s string', 'suffix')
o.parse!(into: m)

if ARGV.length != 1
   o.banner = 'add [flags] <stem>'
   puts o
   exit 1
end

puts m[:p] + ARGV[0] + m[:s]
