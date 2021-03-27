class date:
   month = 12
   day = 31
   def year(self):
      self.month = 1
      self.day = 1

d = date()
d.year()
print(d.day == 1)
