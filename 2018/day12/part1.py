import re

def spread(initial_state, gens, rules, output):
    zero = 30
    state = "." * zero + initial_state + "." * zero
    for x in range(0,gens):
        mods = []
        for i in range(2, len(state)-2):
            current = state[i-2:i+3]
            out = "."
            for j in range(0, len(rules)):
                if rules[j] == current:
                    out = output[j]
                    break
            mods.append([i, out])

        for i, mod in enumerate(mods):
            idx, out = mod
            state = state[:idx] + out + state[idx+1:]

    potIndices = []
    for i, pot in enumerate(state):
        if state[i] == "#":
            potIndices.append(i-zero)
    return sum(potIndices)


def main():
    f = open("input.txt", "r")
    inp = f.readlines()
    initial_state = inp[0].split(":")[1].strip()
    rules = []
    output = []
    for i in range(2, len(inp)):
        r, o = inp[i].replace(" ", "").split("=>")
        rules.append(r.strip())
        output.append(o.strip())
    idx_sum = spread(initial_state, 20, rules, output)
    print(idx_sum)



if __name__ == "__main__":
    main()
