require 'fileutils'

FileUtils.cp("/tmp/foo", "/tmp/bar")
FileUtils.rm("/tmp/foo")
FileUtils.mv("/tmp/bar", "/tmp/foo")
