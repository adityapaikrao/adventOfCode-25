from typing import List 
from itertools import product

def get_accesible_rolls(grid: List[List[str]], max_nbrs: int = 3) -> int:
    N, M = len(grid), len(grid[0])
    num_rolls = 0

    directions = [(di, dj) for di, dj in product([-1, 0, 1], repeat=2) if (di, dj) != (0, 0)]

    for i in range(N):
        for j in range(M):
            if grid[i][j] != "@": continue
            
            num_nbrs = 0
            for x, y in directions:
                
                new_i, new_j = i + x, j + y

                if 0 <= new_i < N and 0 <= new_j < M and grid[new_i][new_j] == "@":
                    num_nbrs += 1

            if num_nbrs <= max_nbrs: 
                num_rolls += 1

    return num_rolls


def parse_input() -> List[List[str]]:
    grid = []

    with open("puzzle1.in", "r") as f:
        for line in f.readlines():
            row = list(line.strip())
            grid.append(row)
        
    return grid

if __name__ == "__main__":
    grid  = parse_input()

    print(get_accesible_rolls(grid, 3))
