require 'getoptlong'
pre = ''
suf = ''

GetoptLong.new(
   ['-p', GetoptLong::REQUIRED_ARGUMENT],
   ['-s', GetoptLong::REQUIRED_ARGUMENT]
).each { |key, val|
   case key
   when '-p'
      pre = val
   when '-s'
      suf = val
   end
}

if ARGV.length != 1
   puts <<eof
add [flags] <stem>
-p string
   prefix
-s string
   suffix
eof
   exit 1
end

puts pre + ARGV[0] + suf
