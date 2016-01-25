package calc

import (
	"strconv"
	"strings"
)

// String set type
type StrSet map[string]struct{}

// Returns a new StrSet.
func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

// BinOp is a binary operator type.
type BinOp func(int, int) int

// PrecMap is a map keyed by operator to set of higher precedence operators.
type PrecMap map[string]StrSet

func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(opMap, prec, expr)
	}
}

// Eval returns the evaluation result of the given expr.
// The expr can have operators defined in opMap and
// decimal integers.
func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
	ops := []string{"("} // 초기 여는 괄호
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				// 더 낮은 순위 연산자이므로 여기서 계산 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			// 닫는 괄호는 여는 괄호까지 계산을 하고 제거
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce(")") // 초기의 여는 괄호까지 모두 계산
	return nums[0]
}
