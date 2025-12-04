from puzzle1 import parse_input
from typing import List, Tuple
from itertools import product
from collections import deque

def is_valid(x: int, y: int) -> bool:
    return 0 <= x < N and 0 <= y < M

def get_num_nbrs(x: int, y: int) -> int:
    
    num_nbrs = 0
    for dx, dy in directions:
        new_x = x + dx
        new_y = y + dy

        if is_valid(new_x, new_y) and grid[new_x][new_y] == "@":
            num_nbrs += 1
    
    return num_nbrs


def get_removable_rolls(max_nbrs: int = 3) -> List[Tuple[int, int]]:
    removable_rolls = []

    for i in range(N):
        for j in range(M):
            if grid[i][j] != "@": continue
        
            num_nbrs = get_num_nbrs(i, j)
            if num_nbrs <= max_nbrs: 
                removable_rolls.append((i, j))
    
    return removable_rolls


def remove_rolls(max_nbrs: int = 3):
    q = deque()
    num_removed = 0

    q.extend(get_removable_rolls(max_nbrs))
    num_removed += len(q)
    print(f"Added {len(q)} roll(s) to be removed")

    while q: 
        for _ in range(len(q)):
            curr_x, curr_y = q.popleft()
            grid[curr_x][curr_y] = "."
        
        q.extend(get_removable_rolls(max_nbrs))
        num_removed += len(q)
        print(f"Added {len(q)} roll(s) to be removed")
    
    return num_removed


if __name__ == "__main__":
    global N, M, directions, grid

    grid = parse_input()
    directions = [(di, dj) for di, dj in product([-1, 0, 1], repeat=2) if (di, dj) != (0, 0)]
    N, M  = len(grid), len(grid[0])

    print(remove_rolls(max_nbrs=3))