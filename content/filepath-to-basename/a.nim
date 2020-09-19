import os
var s = "C:\\Windows\\notepad.exe"
let (s1, s1a, s1b) = splitFile(s)
echo s1 == "C:\\Windows" and s1a == "notepad" and s1b == ".exe"
