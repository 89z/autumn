require 'getoptlong'
pre_s = ''
suf_s = ''

GetoptLong.new(
   ['-p', GetoptLong::REQUIRED_ARGUMENT],
   ['-s', GetoptLong::REQUIRED_ARGUMENT]
).each { |key_s, val_s|
   case key_s
   when '-p'
      pre_s = val_s
   when '-s'
      suf_s = val_s
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

puts pre_s + ARGV[0] + suf_s
