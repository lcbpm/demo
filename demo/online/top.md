### top

### 总体进程信息：

```text

Processes: 585 total, 2 running, 583 sleeping, 2498 threads

Total: 当前系统中总共的进程数量。
Running: 正在运行的进程数量。
Sleeping: 处于睡眠状态的进程数量。
Threads: 当前系统中总共的线程数量。

```
### 系统负载和CPU使用情况

```text

Load Avg: 1.39, 1.52, 1.57  CPU usage: 1.90% user, 2.13% sys, 95.96% idle

Load Avg: 系统负载平均值（1分钟、5分钟和15分钟）。
CPU usage: CPU使用情况，包括用户态使用（user），系统态使用（sys）和空闲时间（idle）。
```

### 内存使用情况：


```text
SharedLibs: 766M resident, 154M data, 131M linkedit.
MemRegions: 134846 total, 5672M resident, 824M private, 3058M shared.
PhysMem: 22G used (2263M wired, 6000K compressor), 1232M unused.

SharedLibs: 共享库的内存使用情况。
MemRegions: 内存区域的统计信息。
PhysMem: 物理内存使用情况，包括已使用的内存和未使用的内存。

```


### 虚拟内存使用情况：

```text

VM: 229T vsize, 4915M framework vsize, 0(0) swapins, 0(0) swapouts.
VM: 虚拟内存使用情况，包含虚拟内存大小，框架虚拟内存大小，以及swap in和swap out的次数。

```

### 网络和磁盘IO：

```text

Networks: packets: 282284/169M in, 270138/94M out.
Disks: 507014/14G read, 179737/3113M written.

Networks: 网络传输情况，包括传入和传出数据包的数量和大小。
Disks: 磁盘IO情况，包括读写操作的次数和数据量。

```

### 进程详细信息：

```text

PID    COMMAND    %CPU    TIME          #TH    #WQ    #POR    MEM    PURG    CMPR    PGRP    PPID
1923   goland     10.4    05:12.58      109    5      678     2932M  143M    0B      1923    1

PID: 进程ID。
COMMAND: 进程名。
%CPU: 该进程的CPU使用率。
TIME: 该进程的累计CPU时间。
#TH: 该进程的线程数量。
#WQ: 该进程的工作队列数量。
#POR: 该进程打开的端口数量。
MEM: 该进程的内存使用量。
PURG: 该进程的可回收内存量。
CMPR: 该进程的压缩内存量。
PGRP: 该进程的进程组ID。
PPID: 该进程的父进程ID。

```

### top -pid <PID> 查看单个进程

