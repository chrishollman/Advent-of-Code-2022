package adventofcode2022

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type opcode string

const (
	multiply opcode = "*"
	add      opcode = "+"
	subtract opcode = "-"
	divide   opcode = "/"
	square   opcode = "^"
)

type operation struct {
	operator opcode
	rhs      *int
}

type monkey struct {
	Items        *Queue[int]
	Operation    *operation
	Test         *operation
	ThrowCounter int
	ThrowTrue    string
	ThrowFalse   string
}

func (m *monkey) takeTurn(trueMonkey *monkey, falseMonkey *monkey, lcm *int) {
	for !m.Items.IsEmpty() {
		m.ThrowCounter++
		worryLevel, _ := m.Items.Dequeue()
		if lcm != nil {
			worryLevel = execOperation(worryLevel, m.Operation) % *lcm
		} else {
			worryLevel = execOperation(worryLevel, m.Operation) / 3
		}
		if execTest(worryLevel, m.Test) {
			trueMonkey.Items.Enqueue(worryLevel)
		} else {
			falseMonkey.Items.Enqueue(worryLevel)
		}
	}
}

func newMonkey(
	items *Queue[int],
	op *operation,
	testCondition *operation,
	throwTargetTrue string,
	throwTargetFalse string,
) *monkey {
	return &monkey{
		Items:      items,
		Operation:  op,
		Test:       testCondition,
		ThrowTrue:  throwTargetTrue,
		ThrowFalse: throwTargetFalse,
	}
}

func execOperation(startingValue int, o *operation) int {
	switch o.operator {
	case multiply:
		return startingValue * (*o.rhs)
	case add:
		return startingValue + (*o.rhs)
	case divide:
		return startingValue / (*o.rhs)
	case subtract:
		return startingValue - startingValue
	}
	return startingValue * startingValue
}

func execTest(startingValue int, o *operation) bool {
	if startingValue%(*o.rhs) != 0 {
		return false
	}
	return true
}

func parseOperation(input string) *operation {
	split := strings.Split(input, " ")

	var op opcode
	switch split[4] {
	case "*":
		op = multiply
	case "+":
		op = add
	case "-":
		op = subtract
	case "/":
		op = divide
	}

	intVal, err := strconv.Atoi(split[5])
	if err != nil {
		return &operation{
			operator: square,
			rhs:      nil,
		}
	}

	return &operation{
		operator: op,
		rhs:      &intVal,
	}
}

func parseTest(input string) *operation {
	split := strings.Split(strings.Trim(input, " "), " ")
	switch split[1] {
	case "divisible":
		intVal, _ := strconv.Atoi(split[len(split)-1])
		return &operation{
			operator: divide,
			rhs:      &intVal,
		}
	default:
		panic("Unsupported test")
	}
}

func monkeyRound(m map[string]*monkey, lcm *int) {
	for i := 0; i < len(m); i++ {
		currMonkey := m[fmt.Sprintf("%v", i)]
		currMonkey.takeTurn(m[currMonkey.ThrowTrue], m[currMonkey.ThrowFalse], lcm)
	}
}

func parseStringIntoQueueOfInts(r *regexp.Regexp, input string) *Queue[int] {
	tmp := r.FindAllString(input, -1)

	var q = Queue[int]{}
	for _, val := range tmp {
		valInt, _ := strconv.Atoi(val)
		q.Enqueue(valInt)
	}

	return &q
}

func dayElevenChallengeOne(input []string) int {

	// Regex for quick number extractions where possible
	rNumbers := regexp.MustCompile("[0-9]{1,}")

	monkeyMap := map[string]*monkey{}
	for i := 0; i < len(input); i = i + 7 {
		identifier := rNumbers.FindString(input[i])
		monkeyMap[identifier] = newMonkey(
			parseStringIntoQueueOfInts(rNumbers, strings.Trim(input[i+1], " ")),
			parseOperation(strings.Trim(input[i+2], " ")),
			parseTest(strings.Trim(input[i+3], " ")),
			rNumbers.FindString(strings.Trim(input[i+4], " ")),
			rNumbers.FindString(strings.Trim(input[i+5], " ")),
		)
	}

	for round := 0; round < 20; round++ {
		monkeyRound(monkeyMap, nil)
	}

	throws := []int{}
	for _, m := range monkeyMap {
		throws = append(throws, m.ThrowCounter)
	}

	sort.Ints(throws)
	return throws[len(throws)-1] * throws[len(throws)-2]
}

func dayElevenChallengeTwo(input []string) int {
	lcm := 1

	// Regex for quick number extractions where possible
	rNumbers := regexp.MustCompile("[0-9]{1,}")

	monkeyMap := map[string]*monkey{}
	for i := 0; i < len(input); i = i + 7 {
		test := parseTest(strings.Trim(input[i+3], " "))
		lcm = lcm * (*test.rhs)

		identifier := rNumbers.FindString(input[i])
		monkeyMap[identifier] = newMonkey(
			parseStringIntoQueueOfInts(rNumbers, strings.Trim(input[i+1], " ")),
			parseOperation(strings.Trim(input[i+2], " ")),
			test,
			rNumbers.FindString(strings.Trim(input[i+4], " ")),
			rNumbers.FindString(strings.Trim(input[i+5], " ")),
		)
	}

	// Run the rounds
	for round := 0; round < 10000; round++ {
		monkeyRound(monkeyMap, &lcm)
	}

	// Find the monkeybusiness score
	throws := []int{}
	for _, m := range monkeyMap {
		throws = append(throws, m.ThrowCounter)
	}

	sort.Ints(throws)
	return throws[len(throws)-1] * throws[len(throws)-2]
}
