package syscall

func ForkOnlyBSDTest() {
	var r1 uintptr
	var pid int
	var err1 Errno
	var wstatus WaitStatus

	ForkLock.Lock()
	runtime_BeforeFork()

	r1, _, err1 = RawSyscall(SYS_FORK, 0, 0, 0)
	if r1 == 0 {
		for {
			RawSyscall(SYS_EXIT, 253, 0, 0)
		}
	}

	runtime_AfterFork()
	ForkLock.Unlock()

	if err1 != 0 || r1 < 1 {
		return
	}

	pid = int(r1)

	_, err2 := Wait4(pid, &wstatus, 0, nil)
	for err2 == EINTR {
		_, err2 = Wait4(pid, &wstatus, 0, nil)
	}

	return
}
