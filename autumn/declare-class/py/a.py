class date:
   month = 1
   day = 1
   def add(self):
      self.day = self.day + 1

d = date()
d.add()
print(d.day == 2)
