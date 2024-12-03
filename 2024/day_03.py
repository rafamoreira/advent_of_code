filename = "day_03.input"

with open(filename) as file:
    data = file.read().splitlines()


data = "".join(data)
print(data)


def procces_line(line: str) -> list[tuple[int, int]]:
    curr = 0
    factors: list[tuple[int, int]] = []
    enabled = True
    while curr < len(line):
        if line[curr : curr + 3] == "mul":
            curr += 3
            if line[curr] == "(":
                curr += 1
                first_factor = ""
                while line[curr].isdigit():
                    first_factor += line[curr]
                    curr += 1
                if len(first_factor) > 3:
                    continue
                if line[curr] == ",":
                    curr += 1
                    second_factor = ""
                    while line[curr].isdigit():
                        second_factor += line[curr]
                        curr += 1
                    if len(second_factor) > 3:
                        continue
                    if line[curr] == ")":
                        curr += 1
                        if enabled:
                            factors.append((int(first_factor), int(second_factor)))
                    else:
                        continue
        elif line[curr : curr + 4] == "do()":
            curr += 4
            enabled = True
        elif line[curr : curr + 7] == "don't()":
            curr += 7
            enabled = False
        else:
            curr += 1

    return factors


factors = procces_line(data)


elements: list[int] = []
for factor in factors:
    elements.append(factor[0] * factor[1])

print(sum(elements))
