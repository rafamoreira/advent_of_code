from copy import deepcopy

filename = "day_02.input"

with open(filename) as file:
    data = file.read().splitlines()

data = [[int(x) for x in y.split(" ")] for y in data]


def is_valid_pair(cur: int, nex: int, increasing: bool | None) -> bool:
    if cur == nex:
        return False
    inc = cur + 3
    dec = cur - 3
    if nex <= inc and nex > cur:
        if increasing:
            return True
        else:
            return False
    elif nex >= dec and nex < cur:
        if not increasing:
            return True
        else:
            return False
    else:
        return False


def validate_report(
    report: list[int], increasing: None | bool = None
) -> tuple[bool, int, bool]:
    i = 0
    while i < len(report) - 1:
        cur = report[i]
        nex = report[i + 1]
        if increasing is None:
            increasing = cur < nex
        if not is_valid_pair(cur, nex, increasing):
            return (False, count, increasing)
        i += 1
    return (True, count, increasing)


count = 0
for report in data:
    safe, i, increasing = validate_report(report)
    if safe:
        count += 1
    else:
        for i in range(len(report)):
            r = deepcopy(report)
            _ = r.pop(i)
            safe, _, _ = validate_report(r)
            if safe:
                count += 1
                break


print("count: ", count)
