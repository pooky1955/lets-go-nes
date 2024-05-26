from bs4 import BeautifulSoup
import pandas as pd
with open("6502 Reference.html","r") as f:
    content = f.read()

bs = BeautifulSoup(content)
bs.prettify()
instruction_names = []
for instruction in bs.select("h3"):
    instruction_names.append(instruction.text.split(' -')[0])

def clean(text : str):
    return text.replace("\n","").strip()
address_map = {
    'Accumulator' : 'ACC',
    'ZeroPage': 'ZP',
    'ZeroPage,X' : 'ZPX',
    'ZeroPage,Y' : 'ZPY',
    'Absolute': 'ABS',
    'Absolute,X' : 'ABX',
    'Implied' : 'IMP',
    'Absolute,Y' : 'ABY',
    '(Indirect,X)' : 'IDX',
    '(Indirect),Y' : 'IDY',
    'Relative' : 'REL',
    'Indirect' : "IND",
    'Immediate' : 'IMM'
}
def process_row(row):
    address_mode, opcode, bytes, cycles = row
    return address_map[address_mode.replace(' ','')],opcode[1:],int(bytes),int(cycles[0])
desired_tables = [table for i,table in enumerate(bs.select("table")) if i >= 2 and i %2 == 0]
info = []
for i, table in enumerate(desired_tables):
    rows = table.select("tr")[1:]
    table_info = [[instruction_names[i]] + list(process_row([clean(td.text) for td in row.select("td")])) for row in rows]
    info.extend(table_info)

print(info)
df = pd.DataFrame(info,columns=['operation','address_mode','opcode','bytes','cycle_count'])
df.to_csv('instructiontable.csv',index=False)