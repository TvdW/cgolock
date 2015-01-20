cgolock is a simple piece of code to prevent Go's *cgo* from spawning hundreds of threads when you really don't want hundreds of threads.

## Usage

Call `cgolock.Lock()` just before calling *cgo*, and `cgolock.Unlock()` when you're done.
