import re
##NOT CORRECT YET
class Machine:
    def call_all(self, inp, out, current):
        outputs = []
        outputs.append(self.addr(inp, out, current))
        outputs.append(self.addi(inp, out, current))
        outputs.append(self.mulr(inp, out, current))
        outputs.append(self.muli(inp, out, current))
        outputs.append(self.banr(inp, out, current))
        outputs.append(self.bani(inp, out, current))
        outputs.append(self.borr(inp, out, current))
        outputs.append(self.bori(inp, out, current))
        outputs.append(self.setr(inp, out, current))
        outputs.append(self.seti(inp, out, current))
        outputs.append(self.gtir(inp, out, current))
        outputs.append(self.gtri(inp, out, current))
        outputs.append(self.gtrr(inp, out, current))
        outputs.append(self.eqir(inp, out, current))
        outputs.append(self.eqri(inp, out, current))
        outputs.append(self.eqrr(inp, out, current))
        return outputs
    def addr(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] + current[inp[1]]
        return output
    def addi(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] + inp[1]
        return output
    def mulr(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] * current[inp[1]]
        return output
    def muli(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] * inp[1]
        return output
    def banr(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] & current[inp[1]]
        return output
    def bani(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] & inp[1]
        return output
    def borr(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] | current[inp[1]]
        return output
    def bori(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]] | inp[1]
        return output
    def setr(self, inp, out, current):
        output = current[:]
        output[out] = current[inp[0]]
        return output
    def seti(self, inp, out, current):
        output = current[:]
        output[out] = inp[0]
        return output
    def gtir(self, inp, out, current):
        output = current[:]
        if inp[0] > current[inp[1]]:
            output[out] = 1
        else:
            output[out] = 0
        return output
    def gtri(self, inp, out, current):
        output = current[:]
        if current[inp[0]] > inp[1]:
            output[out] = 1
        else:
            output[out] = 0
        return output
    def gtrr(self, inp, out, current):
        output = current[:]
        if current[inp[0]] > current[inp[1]]:
            output[out] = 1
        else:
            output[out] = 0
        return output
    def eqir(self, inp, out, current):
        output = current[:]
        if inp[0] == current[inp[1]]:
            output[out] = 1
        else:
            output[out] = 0
        return output
    def eqri(self, inp, out, current):
        output = current[:]
        if current[inp[0]] == inp[1]:
            output[out] = 1
        else:
            output[out] = 0
        return output
    def eqrr(self, inp, out, current):
        output = current[:]
        if current[inp[0]] == current[inp[1]]:
            output[out] = 1
        else:
            output[out] = 0
        return output

##MANUAL LABOR
nstruction_mapping = {
    0: "gtrr",
    1: "mulr",
    2: "eqri",
    3: "addr",
    4: "eqrr",
    5: "addi",
    6: "gtri",
    7: "bori",
    8: "borr",
    9: "muli",
    10: "gtir",
    11: "setr",
    12: "eqir",
    13: "bani",
    14: "seti",
    15: "banr"
}

machine_instructions = [
        "addr",
        "addi",
        "mulr",
        "muli",
        "banr",
        "bani",
        "borr",
        "bori",
        "setr",
        "seti",
        "gtir",
        "gtri",
        "gtrr",
        "eqir",
        "eqri",
        "eqrr",
]

def run_program(codes, instruction_mapping):
    machine = Machine()
    registers = [0, 0, 0, 0]
    for code in codes:
        instr = code[0]
        inp = code[1:3]
        out = code[3]
        func = getattr(machine, machine_instructions[instruction_mapping[instr]])
        registers = func(inp, out, registers)
    print(registers)

def resolve_instruction_mapping(opcodes):
    mapping = {}
    while len(mapping.keys()) != 16:
        for key in opcodes.keys():
            if len(opcodes[key]) == 1:
                actual = opcodes[key][0]
                mapping[key] = actual
                for key in opcodes.keys():
                    if actual in opcodes[key]:
                        opcodes[key].remove(actual)
    print(mapping)
    return mapping

def run(before, instructions, after):
    machine = Machine()
    opcodes = {}
    for i, regs in enumerate(before):
        after_ops = machine.call_all(instructions[i][1:3], instructions[i][3], regs)
        for j, after_op in enumerate(after_ops):
            if after_op == after[i]:
                instr = opcodes.get(j, [])
                if instructions[i][0] not in instr:
                    instr.append(instructions[i][0])
                opcodes[j] = instr
    return opcodes


def main():
    lines = open("lines.txt", "r").readlines()
    program = open("program.txt", "r").readlines()
    regex = re.compile("[0-9]+")
    k = 0
    before = []
    instructions = []
    after = []
    while k < len(lines):
        before.append([int(d) for d in regex.findall(lines[k])])
        instructions.append([int(d) for d in regex.findall(lines[k+1])])
        after.append([int(d) for d in regex.findall(lines[k+2])])
        k += 4
    opcodes = run(before, instructions, after)
    codes = []
    k = 0
    while k < len(program):
        codes.append([int(d) for d in regex.findall(program[k])])
        k += 1
    instruction_mapping = resolve_instruction_mapping(opcodes)
    run_program(codes, instruction_mapping)

if __name__ == "__main__":
    main()
