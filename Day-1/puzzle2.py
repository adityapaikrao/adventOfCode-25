from typing import List, Tuple


def get_password(rotations: List[Tuple[int, int]]) -> int:
    prev_pos = 50
    num_zeroes = 0

    for dir, rotation in rotations:
        # add in complete circles of rotation
        num_zeroes += rotation // 100
        rotation %= 100

        # add in zeroes during rotation
        new_pos = prev_pos + (dir * rotation)
        if (new_pos <= 0 or new_pos >= 100) and prev_pos != 0: 
            num_zeroes += 1

    
        prev_pos = new_pos % 100

    return num_zeroes

if __name__ == "__main__":

    rotations = []
    with open("./puzzle1.in", "r") as f:
        for line in f.readlines():
            dir = 1 if line[0] == "R" else -1
            rotations.append((dir, int(line[1:])))
    
    print(get_password(rotations))