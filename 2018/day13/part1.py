from enum import Enum

class Compass:
    def turn(self, direction, turn):
        if Direction.NORTH:
            return self.handle_north(turn)
        elif Direction.EAST:
            return self.handle_east(turn)
        elif Direction.SOUTH:
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

class Turn(Enum):
    LEFT = 0
    RIGHT = 1
    STRAIGHT = 2

class Direction(Enum):
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
        for i, row in enumerate(grid):
            for j, track in enumerate(row):
                if is_cart(track):
                    cart_index = find_cart(carts, i, j)
                    print(carts[cart_index])
                    next_direction, delta = compass.turn(carts[cart_index].direction, Turn.STRAIGHT)
                    next_i = carts[cart_index].x + delta[0]
                    next_j = carts[cart_index].y + delta[1]
                    if is_intersection(grid[next_i][next_j]):
                        carts[cart_index].x = next_i
                        carts[cart_index].y = next_j
                        direction_to_place, to_place_delta = compass.turn(carts[cart_index].direction, carts[cart_index].get_turn())
                        carts[cart_index].turn = carts[cart_index].turn + 1
                        carts[cart_index].direction = direction_to_place
                    elif is_curve(grid[next_i][next_j]):
                        carts[cart_index].x = next_i
                        carts[cart_index].y = next_j
                        direction_to_place = compass.figure_turn(carts[cart_index].direction, grid[next_i][next_j])
                        carts[cart_index].direction = direction_to_place
                    elif is_cart(grid[next_i][next_j]):
                        return next_j, next_i
                    else:
                        carts[cart_index].x = next_i
                        carts[cart_index].y = next_j
        print_grid(grid)

def print_grid(grid):
    for i, row in enumerate(grid):
        for j, track in enumerate(row):
            print(grid[i][j], end='')
        print()
    print("DONE")

def main():
    f = open("test_input.txt", "r")
    lines = f.readlines()
    grid = []
    carts = []
    compass = Compass()
    for i, line in enumerate(lines):
        line = line.strip()
        grid.append([])
        for j, char in enumerate(line):
            grid[i].append(char)
            if is_cart(char):
                carts.append(Cart(i,j, compass.cart_direction(char)))
    x, y = run_simulation(grid, carts)
    print(x, y)

if __name__ == "__main__":
    main()
