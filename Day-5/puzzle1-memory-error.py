from typing import List, Tuple


def merge_ranges(ranges: List[List[int]]) -> List[List[int]]:
    i = 1
    prev = 0

    while i < len(ranges):
        if ranges[prev][1] >= ranges[i][0]:
            ranges[prev][1] = ranges[i][1]
        else:
            prev += 1
            ranges[prev] = ranges[i]
            
        i += 1
    
    return ranges[:prev + 1]
    
def line_sweep(ranges: List[List[int]], array: List[int]) -> List[int]:
    for start, end in ranges:
        array[start] += 1
        if end + 1 < len(array):
            array[end + 1] -= 1
    
    return array


"""
THIS APPROACH GIVES MEMORY ERROR FOR CREATING THE LONG INGREDIENTS LIST
"""

def get_fresh_ingredients(ranges: List[List[int]], ingredients: List[int]) -> int:
    num_fresh = 0
    ranges.sort() # O(M.log(M))

    ingredients_set = set(ingredients) # O(N)

    merged_ranges = merge_ranges(ranges) # O(M)
    ingredients_list = [0] * (max(ingredients) + 1) # O(N)

    diff_array = line_sweep(merged_ranges, ingredients_list) # O(M)

    cum_sum = 0
    for idx, elem in enumerate(diff_array): # O(M)
        cum_sum += elem

        if idx in ingredients_set and cum_sum > 0:
            num_fresh += 1

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

    print(merge_ranges([[1, 3], [2, 5], [6, 7]]))