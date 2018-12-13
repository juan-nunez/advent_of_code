def grid_power_level(serial_num):
    grid = [[0 for x in range(300)] for y in range(300)]
    for i in range(300):
        for j in range(300):
            grid[j][i] = find_power_level(j, i, serial_num)
    return grid

def form_3x3(grid):
    for i in range(0, 300):
        for j in range(0, 300):
            summation = 0
            for k in range(j, j+3):
                if k == 300:
                    break
                for l in range(i, i+3):
                    if l == 300:
                        break
                    summation += grid[k][l]
            grid[j][i] = summation
    return grid

def best_3x3(grid):
    x = 0
    y = 0
    max_sum = 0
    for i in range(300):
        for j in range(300):
            if grid[j][i] >= max_sum:
                x = j
                y = i
                max_sum = grid[j][i]
    return [x,y]


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
    grid = form_3x3(grid)
    best_x, best_y = best_3x3(grid)
    print(str(best_x) + "," + str(best_y))


if __name__ == "__main__":
    main()
