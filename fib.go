package main

type fibResult struct {
	prev,  curr int
}

func (this *fibResult) next() {
	this.prev, this.curr = this.curr, this.curr + this.prev
}

func fib() *fibResult {
	return &fibResult{1,2}
}
