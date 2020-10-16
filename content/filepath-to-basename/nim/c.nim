import os
var s = r"C:\Windows\notepad.exe"
let (s1, s2, s3) = s.splitFile
echo s1 == r"C:\Windows" and s2 == "notepad" and s3 == ".exe"
