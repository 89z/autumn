var r1 = "a.txt".open
var s1 = r1.readLine
r1.setFilePos(0)
var s2 = r1.readLine
echo [s1, s2]
