import os
let s = r"C:\Windows\notepad.exe"
let (t, u, v) = s.splitFile
echo t == r"C:\Windows" and u == "notepad" and v == ".exe"
