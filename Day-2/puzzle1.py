from typing import List, Tuple


# Part 1
def sum_invalid_ids(ranges: List[Tuple[str, str]]) -> int:
    id_sum = 0

    for interval in ranges:
        for val in range(int(interval[0]), int(interval[1]) + 1):
            val = str(val)
            if len(val) % 2 != 0: continue
            
            mid = len(val) // 2
            if val[:mid] == val[mid:]:
                id_sum += int(val)

    return id_sum


def parse_input(fname: str) -> List[Tuple[str, str]]:
    ranges = []
    with open(f"./{fname}", "r") as f:
        line = f.readline()
        for range in line.split(","):
            range_vals = range.split("-")
            ranges.append((range_vals[0], range_vals[-1]))
    
    return ranges


if __name__ == "__main__":
    ranges = parse_input("puzzle1.in")
    print(sum_invalid_ids(ranges))