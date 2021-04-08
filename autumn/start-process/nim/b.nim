import osproc

let (s, n) = execCmdEx("dust -V")
echo [s == "Dust 0.5.4\n", n == 0]
