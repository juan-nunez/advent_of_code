GRID_SIZE = 301

def grid_power_level(serial_num):
    grid = [[0 for x in range(GRID_SIZE)] for y in range(GRID_SIZE)]
    for i in range(0, GRID_SIZE):
        for j in range(0, GRID_SIZE):
            grid[j][i] = find_power_level(j, i, serial_num)
    return grid

def form_nxn(grid):
    sum_grid = [[0 for x in range(GRID_SIZE)] for y in range(GRID_SIZE)]
    size_grid = [[0 for x in range(GRID_SIZE)] for y in range(GRID_SIZE)]
    max_grid = [[0 for x in range(GRID_SIZE)] for y in range(GRID_SIZE)]

    #create partial sum array
    for i in range(0,GRID_SIZE):
        summation = 0
        for j in range(0, GRID_SIZE):
            summation += grid[j][i]
            sum_grid[j][i] = summation

    for i in range(0, GRID_SIZE):
        summation = 0
        for j in range(0, GRID_SIZE):
            summation += sum_grid[i][j]
            sum_grid[i][j] = summation


    for i in range(1, GRID_SIZE):
        for j in range(1, GRID_SIZE):
            for k in range(1, GRID_SIZE):
                if j+k > GRID_SIZE or i+k > GRID_SIZE:
                    continue
                summation = sum_grid[j+k-1][i+k-1] + sum_grid[j-1][i-1] - sum_grid[j+k-1][i-1] - sum_grid[j-1][i+k-1]
                if summation > max_grid[j][i]:
                    max_grid[j][i] = summation
                    size_grid[j][i] = k
    
    summation = 0
    x = 0
    y = 0
    for i in range(1,GRID_SIZE):
        for j in range(1, GRID_SIZE):
            if max_grid[j][i] > summation:
                print(summation)
                x = j
                y = i
                summation = max_grid[j][i]

    return [x, y, size_grid[x][y]]

#helper func for debugging                
def print_grid(grid):
    for i in range(0, GRID_SIZE):
        for j in range(0, GRID_SIZE):
            print(str(grid[j][i]) + " ", end='')
        print()


def find_power_level(x, y, serial_num):
    rack_id = x + 10
    power_level = rack_id * y
    power_level += serial_num
    power_level *= rack_id
    str_power_level = str(power_level)
    if len(str_power_level) >= 3:
        power_level = int(str_power_level[len(str_power_level)-3])
    else:
        power_level = 0
    return power_level - 5

def main():
    #my puzzle input
    serial_num = 9810
    grid = grid_power_level(serial_num)
    x, y, size = form_nxn(grid)
    print(str(x) + "," + str(y) + "," + str(size))


if __name__ == "__main__":
    main()
