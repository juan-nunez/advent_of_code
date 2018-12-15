import re

def get_repetition(initial_state, gens, rules, output):
    zero = 0
    state = "." * zero + initial_state + "." * zero
    prev = 0
    current = 0
    diffs = []
    for x in range(0,gens):
        mods = []
        if state[0:5] != ".....":
            state = "....." + state
            zero += 5
        if state[:-5] !=  ".....":
            state = state + "....."
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
        current = sum(potIndices)
        # used this function to find out the index where the pattern starts to repeat
        #print(x+1, sum(potIndices), current - prev)
        prev = current
    return 0


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
    idx_sum = get_repetition(initial_state, 200, rules, output)

    print(6739 + 42 * (50000000000 - 159))



if __name__ == "__main__":
    main()
