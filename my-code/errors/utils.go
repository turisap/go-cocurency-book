package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	Result string
	Error  error
}

func computeResult(n int) (res string, err error) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	hasError := rand.Intn(2) == 0

	if hasError {
		err = errors.New("Oh no")
	}

	res = fmt.Sprintf("[w%d]: Computed result from #%d\n", n, n)

	return res, err
}
