from typing import List


def parse_input() -> List[str]:
    banks = []
    with open("puzzle1.in", "r") as f:
        for line in f.readlines():
            banks.append(line.strip())
    
    return banks

# Part 1
def get_total_output(banks: List[str]) -> int:
    total_power = 0

    for bank in banks:
        n = len(bank)
        max_power = 0
        for i in range(n-1):
            for j in range(i+1, n):
                power = int(bank[i]) * 10 + int(bank[j])
                max_power = max(max_power, power)
        
        total_power += max_power
        # print(f"for Bank {bank}:\ntotal_power: {total_power}, max_power: {max_power}\n")

    return total_power




if __name__ == "__main__":
    banks = parse_input()

    print(get_total_output(banks))