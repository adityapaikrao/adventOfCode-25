from typing import List, Tuple


def get_password(rotations: List[Tuple[int, int]]) -> int:
    """
    Implements the logic to solve the second puzzle
    """
    prev_pos = 50
    num_zeroes = 0

    for dir, rotation in rotations:
        # add in complete circles of rotation
        num_zeroes += rotation // 100
        rotation %= 100

        # add in zeroes during rotation and after rotation
        new_pos = prev_pos + (dir * rotation)
        if (new_pos <= 0 or new_pos >= 100) and prev_pos != 0: 
            num_zeroes += 1

    
        prev_pos = new_pos % 100

    return num_zeroes

def parse_input() -> List[Tuple[int, int]]:
    """
    Parses Input from the puzzle.in file into a 2-D List
    """
    rotations = []
    with open("./puzzle1.in", "r") as f:
        for line in f.readlines():
            dir = 1 if line[0] == "R" else -1
            rotations.append((dir, int(line[1:])))
    
    return rotations

if __name__ == "__main__":

    rotations = parse_input()
    
    print(get_password(rotations))