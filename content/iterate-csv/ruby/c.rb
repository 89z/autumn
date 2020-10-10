require 'csv'

s = 'Month,Day
January,Sunday
February,Monday'

CSV.parse(s, headers: true) do |m|
   p m
end
