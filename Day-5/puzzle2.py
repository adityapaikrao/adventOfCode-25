from puzzle1 import parse_input
from typing import List

def merge_ranges(ranges: List[List[int]]) -> List[List[int]]:
    prev = 0
    i = 1

    while i < len(ranges):
        if ranges[prev][1] >= ranges[i][0]:
            ranges[prev][1] = max(ranges[i][1], ranges[prev][1])
        else:
            prev += 1
            ranges[prev] = ranges[i]
        i += 1
    
    return ranges[:prev + 1]

def valid_fresh_ids(ranges: List[List[int]]) -> int:
    num_ids = 0

    ranges.sort()
    merged_ranges = merge_ranges(ranges)

    for start, end in merged_ranges:
        num_ids += end - start + 1
    
    return num_ids

if __name__ == "__main__":
    ranges, _ = parse_input()

    print(valid_fresh_ids(ranges))