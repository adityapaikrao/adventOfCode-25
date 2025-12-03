from typing import List

BANK_LEN = 12
curr_power = []
max_power = 0

"""
########### BACKTRAKCING APPROACH TAKES TOO LONG !!! ###########

def compute_joltage(jolts: List[str]) -> int:
    joltage = 0

    for jolt in jolts:
        joltage = joltage * 10 + int(jolt)
    
    return joltage

def get_max_power(bank: str, idx: int) -> None:
    global curr_power, BANK_LEN, max_power

    # Base Case
    if len(curr_power) == BANK_LEN:
        joltage = compute_joltage(curr_power[:])
        max_power = max(max_power, joltage)
        return 
    elif idx >= len(bank):
        return 


    # Recursive Case
    for i in range(idx, len(bank)):
        curr_power.append(bank[i])

        get_max_power(bank, i + 1)

        curr_power.pop() # backtracking
    
    return 



# Part 2
def get_total_output(banks: List[str]) -> int:
    total_joltage = 0
    global max_power, curr_power

    for bank in banks:
        max_power = 0
        get_max_power(bank, 0) # start at index 0
        
        total_joltage += max_power  
        print(f"for Bank {bank}:\ntotal_power: {total_joltage}, max_power: {max_power}\n")

    return total_joltage
"""
def parse_input() -> List[str]:
    banks = []
    with open("puzzle1.in", "r") as f:
        for line in f.readlines():
            banks.append(line.strip())
    
    return banks

def get_total_output(banks: List[str]) -> int:
    total_joltage = 0

    for bank in banks:
        max_joltage = 0
        n = len(bank)
        start = 0
         

        for l in range(11, -1, -1):
            end = n - l   

            substring = bank[start: end]
            max_digit = max(substring)
            pos = substring.find(max_digit) + start

            
            max_joltage = max_joltage * 10 + int(max_digit)
            # print(start, end, pos, l, max_joltage)
            start = pos + 1

        total_joltage += max_joltage
        print(f"for Bank {bank}:\ntotal_power: {total_joltage}, max_power: {max_joltage}\n")

    return total_joltage    
    


if __name__ == "__main__":
    banks = parse_input()

    print(get_total_output(banks))