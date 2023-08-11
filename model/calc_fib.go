package model

import (
	"math/big"
)

func CalcFib(num int) *big.Int {
	fib_num := big.NewInt(0)
	if num > 0 && num <= 2 {
		fib_num.SetInt64(1)
	} else {
		pre_fib := big.NewInt(1)
		cur_fib := big.NewInt(1)
		tmp_fib := big.NewInt(0)

		for i := 0; i < num-1; i++ {
			tmp_fib.Add(pre_fib, cur_fib)
			pre_fib.Set(cur_fib)
			cur_fib.Set(tmp_fib)
		}
		fib_num.Set(pre_fib)
	}
	return fib_num
}
