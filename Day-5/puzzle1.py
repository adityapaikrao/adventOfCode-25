from typing import List, Tuple
from bisect import bisect_right, bisect_left

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


def get_fresh_ingredients(ranges: List[List[int]], ingredients: List[int]) -> int:
    num_fresh = 0
    # ranges.sort() # O(M.log(M))
    ranges.sort()
    merged_ranges = merge_ranges(ranges)

    ingredients.sort()

    for start, end in merged_ranges:
        left_idx = bisect_left(ingredients, start)
        right_idx = bisect_right(ingredients, end)

        num_fresh += (right_idx - left_idx)

    return num_fresh


def parse_input() -> Tuple[List[List[int]], List[int]]:
    ranges = []
    ingredients = []

    read_ingredients = False
    with open("puzzle1.in", "r") as f:
        for line in f.readlines():
            if line.strip() == "":
                read_ingredients = True
                continue
            
            if read_ingredients:
                ingredients.append(int(line.strip()))
            else:
                range_list = [int(x) for x in line.strip().split("-")]
                ranges.append(range_list)

    return ranges, ingredients


if __name__ == "__main__":
    ranges, ingredients = parse_input()

    print(get_fresh_ingredients(ranges, ingredients)) # O(M log M)

    # print(merge_ranges([[1, 3], [2, 5], [6, 7]]))