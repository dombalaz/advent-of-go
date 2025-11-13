package day10

import (
	"bufio"
	"cmp"
	"io"
	"strconv"
	"strings"
)

// Compares two ordered values and returns them in order as lower and higher.
func lowHigh[T cmp.Ordered](a, b T) (low, high T) {
	if cmp.Less(a, b) {
		return a, b
	}
	return b, a
}

// Reciever is an interface that wraps the behaviour of receiving a microchip value.
type reciever interface {
	receive(n int)
}

// BotBin is a simple storage for two microchip values representing a balance bot.
//
// BotBin can send microchip values to other recievers. It is expected that the 0
// value represent an unset value.
type botBin struct {
	a, b int
}

// Receives microchip value.
//
// It saves n to the first value if unset, otherwise saves to the second value.
func (bin *botBin) receive(n int) {
	if bin.a == 0 {
		bin.a = n
	} else {
		bin.b = n
	}
}

// Sends microchip values to other recievers.
//
// The first reciever would receive the lower of stored values. The function returns false if the botBin does not contain
// both values, in which case the values will not be send to the receivers.
func (bin *botBin) send(l, h reciever) bool {
	if bin.a == 0 || bin.b == 0 {
		return false
	}
	low, high := lowHigh(bin.a, bin.b)
	l.receive(low)
	h.receive(high)
	return true
}

// OutputBin is sa simple storage for a single microchip value.
type outputBin int

// Saves n to the bin.
func (bin *outputBin) receive(n int) {
	*bin = outputBin(n)
}

// valueInstruction is an instruction  that sends a single value to a configured receiver.
type valueInstruction struct {
	done bool
	v    int
	to   reciever
}

// Do the instruction. Always returns true.
func (i *valueInstruction) do() bool {
	if i.done {
		return true
	}
	i.to.receive(i.v)
	i.done = true
	return true
}

// botInstruction is an instruction that takes value from a bot and sends one
// value to the first receiver and the other one to the second receiver.
type botInstruction struct {
	done     bool
	from     *botBin
	to1, to2 reciever
}

// Do the instruction. Returns true if it was able to fully done the instruction, false otherwise.
func (i *botInstruction) do() bool {
	if i.done {
		return true
	}
	ok := i.from.send(i.to1, i.to2)
	if ok {
		i.done = true
	}
	return ok
}

// instruction interface is responsible for running the instruction. Instruction
// is supposed to be run only single time. Returns true if it was done successfuly or if it was already done.
type instruction interface {
	do() bool
}

// instructionSet groups instructions and bins together.
type instructionSet struct {
	instructions []instruction
	botBins      map[int]*botBin
	outputBins   map[int]*outputBin
}

// BotIdWithValues returns a bot if which holds specific values a and b.
//
// Returns true if found, false otherwise.
func (s *instructionSet) BotIdWithValues(a, b int) (int, bool) {
	for k, v := range s.botBins {
		if v.a == a && v.b == b {
			return k, true
		}
		if v.a == b && v.b == a {
			return k, true
		}
	}
	return 0, false
}

// Multiply multiplies values in output bin 0, 1 and 2.
func (s *instructionSet) Multiply() int {
	return int((*s.output(0)) * (*s.output(1)) * (*s.output(2)))
}

// Run iterates over all instructions and performs them until all are done.
func (s *instructionSet) Run() {
	done := false
	for !done {
		done = true
		for _, v := range s.instructions {
			if !v.do() {
				done = false
			}
		}
	}
}

// bot returns bot for gtiven id.
//
// It creates new bot if its missing.
func (s *instructionSet) bot(id int) *botBin {
	bot, ok := s.botBins[id]
	if !ok {
		bot = &botBin{}
		s.botBins[id] = bot
	}
	return bot
}

// output returns output bin for gtiven id.
//
// It creates new output bin if its missing.
func (s *instructionSet) output(id int) *outputBin {
	output, ok := s.outputBins[id]
	if !ok {
		output = new(outputBin)
		s.outputBins[id] = output
	}
	return output
}

// NewInstructions creates new instructionSet from the reader.
func NewInstructions(r io.Reader) (instructionSet, error) {
	set := instructionSet{
		botBins:    make(map[int]*botBin),
		outputBins: make(map[int]*outputBin),
	}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		strs := strings.Split(s.Text(), " ")
		var i instruction
		var err error
		if strs[0][0] == 'v' {
			i, err = set.newValueInstruction(strs)
		} else {
			i, err = set.newBotInstruction(strs)
		}
		if err != nil {
			return set, err
		}
		set.instructions = append(set.instructions, i)
	}
	return set, nil
}

// Creates a single value instruction.
func (s *instructionSet) newValueInstruction(strs []string) (*valueInstruction, error) {
	n, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, err
	}
	to, err := strconv.Atoi(strs[5])
	if err != nil {
		return nil, err
	}
	return &valueInstruction{v: n, to: s.bot(to)}, nil
}

// Creates a single bot instruction.
func (s *instructionSet) newBotInstruction(strs []string) (*botInstruction, error) {
	from, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, err
	}
	to1, err := strconv.Atoi(strs[6])
	if err != nil {
		return nil, err
	}
	to2, err := strconv.Atoi(strs[11])
	if err != nil {
		return nil, err
	}
	var r1, r2 reciever
	if strs[5][0] == 'o' {
		r1 = s.output(to1)
	} else {
		r1 = s.bot(to1)
	}
	if strs[10][0] == 'o' {
		r2 = s.output(to2)
	} else {
		r2 = s.bot(to2)
	}
	bot := s.bot(from)

	return &botInstruction{from: bot, to1: r1, to2: r2}, nil
}
