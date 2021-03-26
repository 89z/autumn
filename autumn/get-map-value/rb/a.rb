# example 1
m = {'month' => 12, 'day' => 31}
n = m['day']
p n == 31

# example 2
m = {month: 12, day: 31}
n = m[:day]
p n == 31
