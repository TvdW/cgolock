// Copyright 2015 Tom van der Woerdt.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgolock

var c chan struct{}

func Init(count int) {
	if c != nil {
		panic("Cannot call cgolock.Init() twice")
	}

	c = make(chan struct{}, count)
}

func Lock() {
	c <- struct{}{}
}

func Unlock() {
	<- c
}
