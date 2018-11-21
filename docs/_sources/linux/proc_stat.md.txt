Whats in /proc/stat
===============

A few people have done a deeper look into this than I will here but this is just a brief
overview to get you up to speed or refresh yourself. I'll link to some much longer posts
about this at the end.

To start with lets see what is even in /proc/stat. Here is the output of running cat against
this file on my laptop running the 4.14.65 linux kernel.

```
eatingthenight@gentoo ~/code/github/bookmarks/source/linux $ cat /proc/stat
cpu  7798454 53129 2254926 19957379 39462 0 6215 0 0 0
cpu0 806091 4325 553814 17936535 37368 0 6215 0 0 0
cpu1 927471 5940 286162 288832 907 0 0 0 0 0
cpu2 1055639 5295 271513 288135 140 0 0 0 0 0
cpu3 1091193 6518 258477 287561 217 0 0 0 0 0
cpu4 1043412 7691 244229 287986 90 0 0 0 0 0
cpu5 976146 8332 220455 288344 485 0 0 0 0 0
cpu6 955370 7495 211697 289516 118 0 0 0 0 0
cpu7 943128 7528 208575 290466 134 0 0 0 0 0
intr 520703432 78 111351 0 0 0 0 0 0 64 395550 0 0 36633262 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 4934 13184 0 111 1737 2674 1960 1643 1415 1657 2055 4795762 1290955 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
ctxt 1097197802
btime 1542001131
processes 1259622
procs_running 2
procs_blocked 0
softirq 308948893 1466144 247063393 238 164781 96944 56 155683 57729916 0 2271738
```

Lets first take a look at the most common section of this which is the first eight lines. This has
been cleaned up and given column names to make it more readable for humans. If you are unfamiliar
with the column names have a look [here][1].

```
     user    nice  system  idle     iowait  irq  softirq steal guest guest_nice
cpu  7798454 53129 2254926 19957379 39462   0    6215    0     0     0
cpu0 806091  4325  553814  17936535 37368   0    6215    0     0     0
cpu1 927471  5940  286162  288832   907     0    0       0     0     0
cpu2 1055639 5295  271513  288135   140     0    0       0     0     0
cpu3 1091193 6518  258477  287561   217     0    0       0     0     0
cpu4 1043412 7691  244229  287986   90      0    0       0     0     0
cpu5 976146  8332  220455  288344   485     0    0       0     0     0
cpu6 955370  7495  211697  289516   118     0    0       0     0     0
cpu7 943128  7528  208575  290466   134     0    0       0     0     0
```

Given the above we can now compute the CPU utlization for a given period of time. To do so you would add up
everything except guest and guest_nice since they are already included in user and nice. Then we add up the
time spent in idle and [iowait][3] to find out how much time the CPU spent doing "nothing". Pretending for
now that I only have one CPU we can calulate the following...

```
>>> total = 7798454 + 53129 + 2254926 + 19957379 + 39462 + 6215
>>> idle = 19957379 + 39462
>>> (total - idle) / total 
0.33586416808080755
```

Which shows that for my first core it's being utlized ~33% of the time. However if I wanted to find my
utilization I could do the following. Which is just a small python script to add up multiple cores and
determine what percent of the time you have not been in the idle state since last time you booted.

```
#!/usr/bin/env python

import re

cpu_info = []

# Read proc stat and extract the lines we care about.
with open("/proc/stat","r") as ps:
  pattern = re.compile("^cpu[0-9]*\s")
  stats = [line.strip() for line in ps if pattern.match(line)]
  for entry in stats:
    cpu_info.append(entry.split()[1:])

user_ticks = 0
nice_ticks = 0
system_ticks = 0
idle_ticks = 0
iowait_ticks = 0
irq_ticks = 0
softirq_ticks = 0
steal_ticks = 0

for entry in cpu_info:
  user_ticks += int(entry[0])
  nice_ticks += int(entry[1])
  system_ticks += int(entry[2])
  idle_ticks += int(entry[3])
  iowait_ticks += int(entry[4])
  irq_ticks += int(entry[5])
  softirq_ticks += int(entry[6])
  steal_ticks += int(entry[7])

# Output utilization since last boot.
total = user_ticks + nice_ticks + system_ticks + idle_ticks + \
  iowait_ticks + irq_ticks + softirq_ticks + steal_ticks

wait = idle_ticks + iowait_ticks

percent = (total - wait) / total

print("Percent Utilized: {}%".format(percent * 100))
```

By running this in a loop you could also determine your utilization over the last X amount of time.
Which would look like so and gives you a nice stream of CPU utilization second by second.

```
#!/usr/bin/env python

import re
import time

# Read proc stat and extract the lines we care about.

last_total = 0
last_wait = 0

while True:
  cpu_info = []

  with open("/proc/stat","r") as ps:
    pattern = re.compile("^cpu[0-9]*\s")
    stats = [line.strip() for line in ps if pattern.match(line)]
    for entry in stats:
      cpu_info.append(entry.split()[1:])
  
  user_ticks = 0
  nice_ticks = 0
  system_ticks = 0
  idle_ticks = 0
  iowait_ticks = 0
  irq_ticks = 0
  softirq_ticks = 0
  steal_ticks = 0
  
  for entry in cpu_info:
    user_ticks += int(entry[0])
    nice_ticks += int(entry[1])
    system_ticks += int(entry[2])
    idle_ticks += int(entry[3])
    iowait_ticks += int(entry[4])
    irq_ticks += int(entry[5])
    softirq_ticks += int(entry[6])
    steal_ticks += int(entry[7])
  
  total = user_ticks + nice_ticks + system_ticks + idle_ticks + \
    iowait_ticks + irq_ticks + softirq_ticks + steal_ticks
  wait = idle_ticks + iowait_ticks

  diff_total = total - last_total
  diff_wait = wait - last_wait
  percent = (diff_total - diff_wait) / diff_total
  
  print("Percent Utilized :{}%".format(percent * 100))
  
  last_total = total
  last_wait = wait
  time.sleep(1)
```


[1]: https://www.opsdash.com/blog/cpu-usage-linux.html
[2]: https://github.com/Leo-G/DevopsWiki/wiki/How-Linux-CPU-Usage-Time-and-Percentage-is-calculated
[3]: https://witekio.com/blog/understanding-io-wait-0-idle-can-ok/
