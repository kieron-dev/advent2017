locks = []
keys = []

with open("input25") as f:
    lines = f.readlines()

i = 0
while i < len(lines):
    l = lines[i].strip()
    if l == "":
        continue
    
    heights = [-1, -1, -1, -1, -1]
    if l == "#####":
        t = locks
    else:
        t = keys
    for j in range(7):
        row = lines[i+j].strip()
        print(row)
        for k in range(5):
            if row[k] == "#":
                heights[k] += 1
    print()
    i += 8
    t.append(heights)


print(locks[0], keys[0])
print(len(locks), len(keys))

count = 0
for lock in locks:
    for key in keys:
        for i in range(5):
            if lock[i] + key[i] > 5:
                fits = False
                break
        else:
            count += 1
print(count)
    
