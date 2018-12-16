import re


class Register:
    def __init__(self, value):
        self.value = value
    def set(self, value):
        self.value = value
    def get(self):
        return self.value


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
        output[out] = current[inp[0]] + inp[1]
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

def run(before, instructions, after):
    registers = [Register(0), Register(0), Register(0), Register(0)]
    machine = Machine()
    total_count_of_gte_3 = 0
    for i, regs in enumerate(before):
        after_ops = machine.call_all(instructions[i][1:3], instructions[i][3], regs)
        cnt = 0
        for after_op in after_ops:
            if after_op == after[i]:
                cnt += 1
        if cnt >= 3:
            total_count_of_gte_3 += 1
    print(total_count_of_gte_3)



def main():
    f = open("input.txt", "r")
    lines = f.readlines()
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
    run(before, instructions, after)

if __name__ == "__main__":
    main()
