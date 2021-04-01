from pathlib import Path
p = Path(r'C:\Windows\notepad.exe')
s = str(p.parent)
print(s == r'C:\Windows')
