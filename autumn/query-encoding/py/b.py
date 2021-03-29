from urllib import parse
m = {'west': 'left', 'east': 'right'}
s = parse.urlencode(m)
print(s == 'west=left&east=right')
