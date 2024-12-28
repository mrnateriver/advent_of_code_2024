package day17

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"mrnateriver.io/advent_of_code_2024/shared"
)

func readInput() cpu {
	cpu := cpu{}

	registerRegex := regexp.MustCompile("^Register ([ABC]): (\\d+)")
	programRegex := regexp.MustCompile("Program: ([\\d,]+)")

	for line := range shared.ReadInputLines("day17/input") {
		matches := registerRegex.FindStringSubmatch(line)
		if matches != nil {
			regName := matches[1]
			value, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(fmt.Errorf("failed to parse register value %v: %w", matches[2], err))
			}

			*registerByName(&cpu, regName) = value
			continue
		}

		matches = programRegex.FindStringSubmatch(line)
		if matches != nil {
			programOpcodes := matches[1]
			split := strings.Split(programOpcodes, ",")
			for _, op := range split {
				intOp, err := strconv.Atoi(op)
				if err != nil {
					panic(fmt.Errorf("failed to parse program opcode %v: %w", op, err))
				}
				cpu.program = append(cpu.program, intOp)
			}
			break
		}
	}

	return cpu
}

func debug(c *cpu) {
	debugPrint(c)
	time.Sleep(30 * time.Millisecond)

	shared.SetCursorColumn(1)
	shared.MoveCursorUp(6)
}

func debugPrint(c *cpu) {
	fill := strings.Repeat(" ", 32)
	fmt.Printf("Register A: %v%s\n", c.regA, fill)
	fmt.Printf("Register B: %v%s\n", c.regB, fill)
	fmt.Printf("Register C: %v%s\n", c.regC, fill)
	fmt.Println(fill, fill, fill)

	var pb strings.Builder
	for _, op := range c.program {
		pb.WriteString(strconv.Itoa(op))
		pb.WriteByte(',')
	}

	fmt.Printf("Program: %v\n", pb.String())
	fmt.Printf("         ")
	for i := 0; i < c.ptr; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf(shared.Colored(shared.ColorBlue, "^ ^%s\n"), fill)
}

type cpu struct {
	regA, regB, regC int
	program          []int
	ptr              int
}

func registerByName(c *cpu, regName string) *int {
	switch regName {
	case "A":
		return &c.regA
	case "B":
		return &c.regB
	case "C":
		return &c.regC
	}
	panic(fmt.Errorf("invalid register requested %v", regName))
}

var comboRegisters = []string{"A", "B", "C"}

func combo(c *cpu, operand int) int {
	if operand >= 7 || operand < 0 {
		panic(fmt.Errorf("invalid combo operand %v", operand))
	}

	if operand < 4 {
		return operand
	}

	return *registerByName(c, comboRegisters[operand-4])
}

type opcode func(*cpu, int) string

var opcodes = []opcode{
	adv,
	bxl,
	bst,
	jnz,
	bxc,
	out,
	bdv,
	cdv,
}

func run(c *cpu, d bool) string {
	outputs := make([]string, 0, 4)

	pl := len(c.program)
	for c.ptr+1 < pl {
		opcode := c.program[c.ptr]
		operand := c.program[c.ptr+1]

		handler := opcodes[opcode]

		output := handler(c, operand)
		if output != "" {
			outputs = append(outputs, output)
		}

		if d {

			debug(c)
		}
	}

	return strings.Join(outputs, ",")
}

// The adv instruction (opcode 0) performs division. The numerator is the value in the A register. The denominator is found by raising 2 to the power of the instruction's combo operand. (So, an operand of 2 would divide A by 4 (2^2); an operand of 5 would divide A by 2^B.) The result of the division operation is truncated to an integer and then written to the A register.
func adv(c *cpu, operand int) string {
	c.regA = c.regA >> combo(c, operand)
	c.ptr += 2
	return ""
}

// The bxl instruction (opcode 1) calculates the bitwise XOR of register B and the instruction's literal operand, then stores the result in register B.
func bxl(c *cpu, operand int) string {
	c.regB = c.regB ^ operand
	c.ptr += 2
	return ""
}

// The bst instruction (opcode 2) calculates the value of its combo operand modulo 8 (thereby keeping only its lowest 3 bits), then writes that value to the B register.
func bst(c *cpu, operand int) string {
	c.regB = combo(c, operand) % 8
	c.ptr += 2
	return ""
}

// The jnz instruction (opcode 3) does nothing if the A register is 0. However, if the A register is not zero, it jumps by setting the instruction pointer to the value of its literal operand; if this instruction jumps, the instruction pointer is not increased by 2 after this instruction.
func jnz(c *cpu, operand int) string {
	if c.regA == 0 {
		c.ptr++
		return ""
	}

	c.ptr = operand
	return ""
}

// The bxc instruction (opcode 4) calculates the bitwise XOR of register B and register C, then stores the result in register B. (For legacy reasons, this instruction reads an operand but ignores it.)
func bxc(c *cpu, operand int) string {
	c.regB = c.regB ^ c.regC
	c.ptr += 2
	return ""
}

// The out instruction (opcode 5) calculates the value of its combo operand modulo 8, then outputs that value. (If a program outputs multiple values, they are separated by commas.)
func out(c *cpu, operand int) string {
	c.ptr += 2
	return strconv.Itoa(combo(c, operand) % 8)
}

// The bdv instruction (opcode 6) works exactly like the adv instruction except that the result is stored in the B register. (The numerator is still read from the A register.)
func bdv(c *cpu, operand int) string {
	c.regB = c.regA >> combo(c, operand)
	c.ptr += 2
	return ""
}

// The cdv instruction (opcode 7) works exactly like the adv instruction except that the result is stored in the C register. (The numerator is still read from the A register.)
func cdv(c *cpu, operand int) string {
	c.regC = c.regA >> combo(c, operand)
	c.ptr += 2
	return ""
}
