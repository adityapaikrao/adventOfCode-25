from typing import List, Tuple


# Part 2
def sum_invalid_ids(ranges: List[Tuple[str, str]]) -> int:
    id_sum = 0

    for interval in ranges:
        curr_sum = 0
        # vals_added = []
        for val in range(int(interval[0]), int(interval[1]) + 1):
            val = str(val)
            
            n = len(val)
            for i in range(1, n // 2 + 1):
                if val[:i] * (n // i) == val:
                    # vals_added.append(val)
                    curr_sum += int(val)
                    break

        id_sum += curr_sum
        # print(f"After interval: {interval}\ncurr_sum: {curr_sum} & id_sum: {id_sum} & vals: {vals_added}\n")
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