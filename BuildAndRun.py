import os
import subprocess

# Update to the latest version
for line in os.popen('git fetch -p -q; git merge -q origin/master').readlines():
    print line.strip()
    
# Move the old version over
for line in os.popen('cp monitor oldmonitor').readlines():
    print line.strip()

# Rebuild
for line in os.popen('go build ./...').readlines():
    print line.strip()

# Rebuild
for line in os.popen('go build').readlines():
    print line.strip()


size_1 = os.path.getsize('./oldmonitor')
size_2 = os.path.getsize('./monitor')

if size_1 != size_2:
    for line in os.popen('killall monitor').readlines():
        pass
    subprocess.Popen('./monitor')
