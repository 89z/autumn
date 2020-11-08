require 'csv'

s = 'Month,Day
January,Sunday
February,Monday'

CSV.new(s, headers: true).each do |m|
   p m
end
