import os, stat
os.access('/etc/shells', os.F_OK)
os.path.exists('/etc/shells')
os.path.isfile('/etc/shells')
stat.S_IFMT(os.stat('/etc/shells').st_mode) == stat.S_IFREG
stat.S_ISREG(os.stat('/etc/shells').st_mode)
