from enum import IntEnum

class Compass:
    def turn(self, direction, turn):
        if direction == Direction.NORTH:
            return self.handle_north(turn)
        elif direction == Direction.EAST:
            return self.handle_east(turn)
        elif direction == Direction.SOUTH:
            return self.handle_south(turn)
        else:
            return self.handle_west(turn)

    def cart_direction(self, cart):
        if cart == '^':
            return Direction.NORTH
        elif cart == '>':
            return Direction.EAST
        elif cart =='v':
            return Direction.SOUTH
        else:
            return Direction.WEST

    def figure_turn(self, direction, curve):
        if direction == Direction.NORTH and curve == '\\':
            return Turn.LEFT
        elif direction == Direction.NORTH and curve == '/':
            return Turn.RIGHT
        elif direction == Direction.EAST and curve == '\\':
            return Turn.RIGHT
        elif direction == Direction.EAST and curve == '/':
            return Turn.LEFT
        elif direction == Direction.SOUTH and curve == '\\':
            return Turn.LEFT
        elif direction == Direction.SOUTH and curve == '/':
            return Turn.RIGHT
        elif direction == Direction.WEST and curve == '\\':
            return Turn.RIGHT
        elif direction == Direction.WEST and curve == '/':
            return Turn.LEFT

    def handle_north(self, turn):
        if turn == Turn.LEFT:
            return [Direction.WEST, (0, -1)]
        elif turn == Turn.RIGHT:
            return [Direction.EAST, (0, 1)]
        else:
            return [Direction.NORTH, (-1,0)]

    def handle_east(self, turn):
        if turn == Turn.LEFT:
            return [Direction.NORTH, (-1, 0)]
        elif turn == Turn.RIGHT:
            return [Direction.SOUTH, (1, 0)]
        else:
            return [Direction.EAST, (0, 1)]

    def handle_south(self, turn):
        if turn == Turn.LEFT:
            return [Direction.EAST, (0, 1)]
        elif turn == Turn.RIGHT:
            return [Direction.WEST, (0, -1)]
        else:
            return [Direction.SOUTH, (1, 0)]

    def handle_west(self, turn):
        if turn == Turn.LEFT:
            return [Direction.SOUTH, (1,0)]
        elif turn == Turn.RIGHT:
            return [Direction.NORTH, (-1, 0)]
        else:
            return [Direction.WEST, (0, -1)]

class Turn(IntEnum):
    LEFT = 0
    STRAIGHT = 1
    RIGHT = 2

class Direction(IntEnum):
    NORTH = 0
    EAST = 1
    SOUTH = 2
    WEST = 3

class Cart:
    def __init__(self, x, y, direction):
        self.x = x
        self.y = y
        self.direction = direction
        self.turn = 0
        self.on_intersection = False
        self.on_curve = False

    def get_turn(self):
        return self.turn % 3

def is_cart(symbol):
    if symbol == '<' or symbol == '>' or symbol == '^' or symbol == 'v':
        return True
    return False

def is_intersection(symbol):
    return symbol == '+'

def is_curve(symbol):
    return symbol == '\\' or symbol == '/'

def find_cart(carts, i, j):
    for k, cart in enumerate(carts):
        if cart.x == i and cart.y == j:
            return k
    return -1


def run_simulation(grid, carts):
    compass = Compass()
    has_crashed = False
    while not has_crashed:
        #print_grid(grid, carts)
        used_carts = []
        for i, row in enumerate(grid):
            for j, track in enumerate(row):
                cart_index = find_cart(carts, i, j)
                if cart_index != -1 and carts[cart_index] not in used_carts:
                    cart = carts[cart_index]
                    used_carts.append(cart)
                    next_direction, delta = compass.turn(cart.direction, Turn.STRAIGHT)
                    next_i = cart.x + delta[0]
                    next_j = cart.y + delta[1]
                    if find_cart(carts, next_i, next_j) != -1:
                        carts.remove(cart)
                        carts.remove(carts[find_cart(carts, next_i, next_j)])
                        print("CAR CRASH", next_i, next_j)
                        print("CARTS LEFT", len(carts))
                    elif is_intersection(grid[next_i][next_j]):
                        cart.x = next_i
                        cart.y = next_j
                        next_direction, to_place_delta = compass.turn(cart.direction, cart.get_turn())
                        cart.turn = carts[cart_index].turn + 1
                        cart.direction = next_direction
                    elif is_curve(grid[next_i][next_j]):
                        cart.x = next_i
                        cart.y = next_j
                        required_turn = compass.figure_turn(cart.direction, grid[next_i][next_j])
                        next_direction, to_place_delta = compass.turn(cart.direction, required_turn)
                        cart.direction = next_direction 
                    else:
                        cart.x = next_i
                        cart.y = next_j
        if len(carts) == 1:
           return carts[0].y, carts[0].x
    return 0,0

def print_grid(grid, carts):
    for i, row in enumerate(grid):
        for j, track in enumerate(row):
            if find_cart(carts, i, j) != -1:
                print("*", end='')
            else: print(grid[i][j], end='')
        print()
    print("DONE")

def main():
    f = open("input.txt", "r")
    lines = f.readlines()
    grid = []
    carts = []
    compass = Compass()
    for i, line in enumerate(lines):
        grid.append([])
        for j, char in enumerate(line):
            if is_cart(char):
                if char == "^" or char == "v":
                    grid[i].append("|")
                else:
                    grid[i].append("-")
                carts.append(Cart(i,j, compass.cart_direction(char)))
            else:
                grid[i].append(char)
    x, y = run_simulation(grid, carts)
    print(x, y)

if __name__ == "__main__":
    main()
