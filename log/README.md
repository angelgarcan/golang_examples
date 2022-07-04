
# Klog example
### Klog Flags
https://pkg.go.dev/k8s.io/klog/v2#section-documentation  

```
	-logtostderr=true
		Logs are written to standard error instead of to files.
             This shortcuts most of the usual output routing:
             -alsologtostderr, -stderrthreshold and -log_dir have no
             effect and output redirection at runtime with SetOutput is
             ignored.
	-alsologtostderr=false
		Logs are written to standard error as well as to files.
	-stderrthreshold=ERROR
		Log events at or above this severity are logged to standard
		error as well as to files.
	-log_dir=""
		Log files will be written to this directory instead of the
		default temporary directory.

	Other flags provide aids to debugging.

	-log_backtrace_at=""
		When set to a file and line number holding a logging statement,
		such as
			-log_backtrace_at=gopherflakes.go:234
		a stack trace will be written to the Info log whenever execution
		hits that statement. (Unlike with -vmodule, the ".go" must be
		present.)
	-v=0
		Enable V-leveled logging at the specified level.
	-vmodule=""
		The syntax of the argument is a comma-separated list of pattern=N,
		where pattern is a literal file name (minus the ".go" suffix) or
		"glob" pattern and N is a V level. For instance,
			-vmodule=gopher*=3
		sets the V level to 3 in all Go files whose names begin "gopher".
```

https://github.com/kubernetes/klog  