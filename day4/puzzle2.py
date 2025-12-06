from puzzle1 import parse_input
from typing import List, Tuple
from itertools import product
from collections import deque

def is_valid(x: int, y: int, N: int, M: int) -> bool:
    return 0 <= x < N and 0 <= y < M

def get_num_nbrs(grid: List[List[str]], x: int, y: int, directions: List[Tuple[int, int]], N: int, M: int) -> int:
    num_nbrs = 0
    for dx, dy in directions:
        new_x = x + dx
        new_y = y + dy
        if is_valid(new_x, new_y, N, M) and grid[new_x][new_y] == "@":
            num_nbrs += 1
    return num_nbrs

def get_removable_rolls(grid: List[List[str]], directions: List[Tuple[int, int]], N: int, M: int, max_nbrs: int = 3) -> List[Tuple[int, int]]:
    removable_rolls = []
    for i in range(N):
        for j in range(M):
            if grid[i][j] != "@":
                continue
            if get_num_nbrs(grid, i, j, directions, N, M) <= max_nbrs:
                removable_rolls.append((i, j))
    return removable_rolls

def remove_rolls(grid: List[List[str]], max_nbrs: int = 3) -> int:
    N, M = len(grid), len(grid[0])
    directions = [(di, dj) for di, dj in product([-1, 0, 1], repeat=2) if (di, dj) != (0, 0)]

    q = deque()
    num_removed = 0

    q.extend(get_removable_rolls(grid, directions, N, M, max_nbrs))
    num_removed += len(q)
    print(f"Added {len(q)} roll(s) to be removed")

    while q:
        for _ in range(len(q)):
            curr_x, curr_y = q.popleft()
            grid[curr_x][curr_y] = "."
        q.extend(get_removable_rolls(grid, directions, N, M, max_nbrs))
        num_removed += len(q)
        print(f"Added {len(q)} roll(s) to be removed")

    return num_removed

if __name__ == "__main__":
    grid = parse_input()
    print(remove_rolls(grid, max_nbrs=3))