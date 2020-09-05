import shutil

shutil.copy('/tmp/foo', '/tmp/bar')
os.remove('/tmp/foo')
shutil.move('/tmp/bar', '/tmp/foo')
