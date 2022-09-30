import pandas as pd

df = pd.read_csv('students-ids.csv')
ids = []
with open('results.txt') as f:
    lines = f.readlines()
    for line in lines:
        ids.append(int(line[:-1]))

res = df.loc[df['id'].isin(ids)]
print(res)