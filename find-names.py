import pandas as pd

df = pd.read_csv('students-ids.csv')
ids = list(map(int, input().split(" ")))

print(df.loc[df['id'].isin(ids)].to_string())