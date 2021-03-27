class Time:
   hours = 23
   def duration(self, minutes):
      return self.hours * 60 + minutes

t = Time()
n = t.duration(59)
print(n == 1439)
