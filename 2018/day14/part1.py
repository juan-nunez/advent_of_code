
def run(num_recipes):
    recipes = [3,7]
    first = 0
    second = 1
    while len(recipes) <= num_recipes:
        n = recipes[first] + recipes[second]
        if n < 10:
            recipes.append(n)
        else:
            dig1, dig2 = [int(d) for d in str(n)]
            recipes.extend([dig1, dig2])
        first = (first + 1 + recipes[first]) % len(recipes)
        second = (second + 1 + recipes[second]) % len(recipes)
    length = len(recipes)
    return "".join(map(str, recipes[num_recipes-10:num_recipes]))


def main():
    num_recipes = 554401 + 10
    print(run(num_recipes))

if __name__ == "__main__":
    main()
