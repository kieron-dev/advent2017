count = 0
with open("input01") as file:
    pos = 50
    for line in file:
        dir = 1
        if line[0] == "L":
            dir = -1
        pos += dir * int(line[1:])
        pos %= 100
        if pos == 0:
            count += 1

print("part 1:", count)

count = 0
with open("input01") as file:
    pos = 50
    for line in file:
        dir = 1
        if line[0] == "L":
            dir = -1
        pos += dir * int(line[1:])
        pos %= 100
        if pos == 0:
            count += 1

print("part 2:", count)
