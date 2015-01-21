cgolock is a simple piece of code to prevent Go's *cgo* from spawning hundreds of threads when you really don't want hundreds of threads. It does this across multiple packages, to ensure that using 20 libraries doesn't still give you `GOMAXPROCS*20` threads.

## Usage

Call `cgolock.Lock()` just before calling *cgo*, and `cgolock.Unlock()` when you're done.
