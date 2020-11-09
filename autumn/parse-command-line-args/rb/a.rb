require 'getoptlong'
s_pre = ''
s_suf = ''

GetoptLong.new(
   ['-p', GetoptLong::REQUIRED_ARGUMENT],
   ['-s', GetoptLong::REQUIRED_ARGUMENT]
).each { |s_key, s_val|
   case s_key
   when '-p'
      s_pre = s_val
   when '-s'
      s_suf = s_val
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

puts s_pre + ARGV[0] + s_suf
