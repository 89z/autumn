unless system("ls -l /tmp")
  raise "ls failed"
end
