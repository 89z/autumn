import datetime, email.utils, time
m1 = datetime.date
m2 = datetime.datetime
m3 = email.utils
s1 = m1.today().isoformat()
s2 = m2.now().isoformat()
s3 = m2.now().astimezone().isoformat()
s4 = m3.formatdate(localtime = 1)
s5 = time.asctime()
t.strftime('%Y-%m-%d %H:%M:%S')
