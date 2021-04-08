import osproc

let s = execProcess("dust -V")
echo s == "Dust 0.5.4\n"
