def run(inp):
    recipes = [3,7]
    first = 0
    second = 1
    #Chose a big number to iterate over
    cnt = 100000000
    for i in range(0, cnt):
        n = recipes[first] + recipes[second]
        if n < 10:
            recipes.append(n)
        else:
            dig1, dig2 = [int(d) for d in str(n)]
            recipes.extend([dig1, dig2])
        first = (first + 1 + recipes[first]) % len(recipes)
        second = (second + 1 + recipes[second]) % len(recipes)

    for i in range(0, cnt-6):
        k = int("".join(map(str, recipes[i:i+6])))
        if k == inp:
            return i


def main():
    inp = 554401
    print(run(inp))

if __name__ == "__main__":
    main()
