require 'optparse'
m = {p: '', s: ''}
p = OptionParser.new
p.on('-p string', 'prefix')
p.on('-s string', 'suffix')
p.parse!(into: m)

if ARGV.length != 1
   p.banner = 'add [flags] <stem>'
   puts p
   exit 1
end

puts m[:p] + ARGV[0] + m[:s]
