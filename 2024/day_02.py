filename = "day_02.input"

with open(filename) as file:
    data = file.read().splitlines()

data = [[int(x) for x in y.split(" ")] for y in data]


def validate_report(report: list[int]) -> bool:
    increasing = None
    while len(report) > 1:
        cur = report.pop(0)
        nex = report[0]
        if cur == nex:
            return False
        inc = cur + 3
        dec = cur - 3
        if nex <= inc and nex > cur:
            if increasing is None or increasing:
                increasing = True
            else:
                return False
        elif nex >= dec and nex < cur:
            if increasing is None or not increasing:
                increasing = False
            else:
                return False
        else:
            return False
    return True


count = 0
for report in data:
    safe = validate_report(report)
    if safe:
        count += 1

print("count: ", count)
