segfaults!

```shell
$ go run -ldflags '-extldflags "-static"' .
fatal error: unexpected signal during runtime execution
[signal 0xb code=0x1 addr=0xe5 pc=0x7fa688e17a5c]

runtime stack:
runtime.throw(0x6e1bc0, 0x2a)
	/usr/local/go/src/runtime/panic.go:527 +0x90
runtime.sigpanic()
	/usr/local/go/src/runtime/sigpanic_unix.go:12 +0x5a

goroutine 5 [syscall, locked to thread]:
runtime.cgocall(0x402170, 0xc82004dcd0, 0xc800000000)
	/usr/local/go/src/runtime/cgocall.go:120 +0x11b fp=0xc82004dc80 sp=0xc82004dc50
os/user._Cfunc_mygetpwuid_r(0x0, 0xc820017620, 0x1a7ef90, 0x400, 0xc82002a0e0, 0x0)
	??:0 +0x39 fp=0xc82004dcd0 sp=0xc82004dc80
os/user.lookupUnix(0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
	/usr/local/go/src/os/user/lookup_unix.go:99 +0x723 fp=0xc82004de40 sp=0xc82004dcd0
os/user.current(0x0, 0x0, 0x0)
	/usr/local/go/src/os/user/lookup_unix.go:39 +0x42 fp=0xc82004de80 sp=0xc82004de40
os/user.Current(0x6a30e0, 0x0, 0x0)
	/usr/local/go/src/os/user/lookup.go:9 +0x24 fp=0xc82004dea0 sp=0xc82004de80
github.com/cockroachdb/cgo_static_boom.TestBoom(0xc820088090)
	/go/src/github.com/cockroachdb/cgo_static_boom/boom_test.go:12 +0xaf fp=0xc82004df58 sp=0xc82004dea0
testing.tRunner(0xc820088090, 0x9c0e30)
	/usr/local/go/src/testing/testing.go:456 +0x98 fp=0xc82004df90 sp=0xc82004df58
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1696 +0x1 fp=0xc82004df98 sp=0xc82004df90
created by testing.RunTests
	/usr/local/go/src/testing/testing.go:561 +0x86d

goroutine 1 [chan receive]:
testing.RunTests(0x6f8608, 0x9c0e30, 0x1, 0x1, 0x1)
	/usr/local/go/src/testing/testing.go:562 +0x8ad
testing.(*M).Run(0xc82004bef8, 0x0)
	/usr/local/go/src/testing/testing.go:494 +0x70
main.main()
	github.com/cockroachdb/cgo_static_boom/_test/_testmain.go:54 +0x116

goroutine 17 [syscall, locked to thread]:
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1696 +0x1
exit status 2
FAIL	github.com/cockroachdb/cgo_static_boom	0.010s
```
