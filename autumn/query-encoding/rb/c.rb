require 'uri'
m = {'month' => 'May', 'day' => 'Friday'}
s = URI.encode_www_form(m)
puts s
