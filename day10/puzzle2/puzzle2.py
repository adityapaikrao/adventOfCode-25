#!/usr/bin/env python3
"""
Solve the joltage configuration puzzle using Integer Linear Programming.

For each machine:
- Let x[i] = number of times to press button i
- Constraint: for each counter j, sum(x[i] * A[i][j]) = target[j]
  where A[i][j] = 1 if button i affects counter j, else 0
- Minimize: sum(x[i])
- All x[i] >= 0 and are integers
"""

from scipy.optimize import milp, LinearConstraint, Bounds
import numpy as np
import re


def parse_input(filepath):
    """Parse the puzzle input file."""
    machines = []
    
    with open(filepath, 'r') as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            
            parts = line.split()
            if len(parts) < 2:
                continue
            
            # Parse joltage requirements (last part in {})
            joltage_str = parts[-1]
            joltage_match = re.findall(r'\{([^}]+)\}', joltage_str)
            if not joltage_match:
                continue
            joltage = list(map(int, joltage_match[0].split(',')))
            
            # Parse buttons (parts between first and last, in parentheses)
            buttons = []
            for i in range(1, len(parts) - 1):
                button_str = parts[i].strip('()')
                if button_str:
                    button = list(map(int, button_str.split(',')))
                else:
                    button = []
                buttons.append(button)
            
            machines.append((buttons, joltage))
    
    return machines


def solve_machine(buttons, joltage):
    """
    Solve for minimum button presses using Integer Linear Programming.
    
    Variables: x[i] = number of times to press button i
    Objective: minimize sum(x[i])
    Constraints: for each counter j, sum(x[i] * A[i][j]) = target[j]
    """
    num_buttons = len(buttons)
    num_counters = len(joltage)
    
    if num_buttons == 0:
        if all(j == 0 for j in joltage):
            return 0
        return None  # No solution
    
    # Build constraint matrix A where A[j][i] = 1 if button i affects counter j
    A = np.zeros((num_counters, num_buttons), dtype=float)
    for i, button in enumerate(buttons):
        for idx in button:
            if idx < num_counters:
                A[idx][i] = 1
    
    # Objective: minimize sum of x[i] (all coefficients are 1)
    c = np.ones(num_buttons)
    
    # Equality constraints: A @ x = joltage
    b = np.array(joltage, dtype=float)
    
    # All variables must be non-negative integers
    bounds = Bounds(lb=0, ub=np.inf)
    
    # Linear constraint: A @ x = joltage (equality means lb = ub = b)
    constraints = LinearConstraint(A, lb=b, ub=b)  # type: ignore[arg-type]
    
    # All variables are integers
    integrality = np.ones(num_buttons)  # 1 means integer
    
    # Solve
    result = milp(c, constraints=constraints, bounds=bounds, integrality=integrality)
    
    if result.success:
        total_presses = int(round(result.fun))
        return total_presses
    else:
        return None


def main():
    filepath = '../puzzle.in'
    machines = parse_input(filepath)
    
    total = 0
    for i, (buttons, joltage) in enumerate(machines):
        result = solve_machine(buttons, joltage)
        if result is not None:
            print(f"{result} presses for target: {joltage}")
            total += result
        else:
            print(f"No solution for target: {joltage}")
    
    print(f"numButtonPresses: {total}")


if __name__ == '__main__':
    main()
