filename = 'day_01.input'

with open(filename) as file:
    data = file.read().splitlines()


data = [
    [y.strip() for y in x.split(' ', 1)] 
    for x in data 
]

separate_data: dict[str, list[int]] = {"left": [], "right": [], "diff": [], "simmilarity": []}

for x, y in data:
    separate_data["left"].append(int(x))
    separate_data["right"].append(int(y))

separate_data["left"] = sorted(separate_data["left"])
separate_data["right"] = sorted(separate_data["right"])

for left, right in zip(separate_data["left"], separate_data["right"]):
    separate_data["diff"].append(abs(left - right))

print("diff: ", sum(separate_data["diff"]))

for left in separate_data["left"]:
    separate_data["simmilarity"].append(left * separate_data["right"].count(left))

print("simmilarity: ", sum(separate_data["simmilarity"]))
