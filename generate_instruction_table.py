import pandas as pd

df = pd.read_csv('instructiontable.csv')
preface = '''
package main
type Instruction struct {
    Opcode int
    Address string
    Cycles int
    Operation string


}
'''
all_operations = sorted(set((df['operation'].values)))
all_addresses = sorted(set((df['address_mode'].values)))
newline = '\n'
tab = '\t'
const_declaration = f'''
const (
{newline.join([tab+f'OP_{op} = "{op}"' for op in all_operations])}
{newline.join([tab+f'ADR_{addr} = "{addr}"' for addr in all_addresses])}
)
'''

instruction_list = []
for _, row in df.iterrows():
    operation,address,opcode,_, cycles = row
    format_string = f'\t0x{opcode} : Instruction{{Opcode : 0x{opcode}, Address : ADR_{address}, Operation: OP_{operation}, Cycles: {cycles}}},'
    instruction_list.append(format_string)

instruction_map_declaration = f'''
var instructionMap map[int]Instruction = map[int]Instruction{{
{newline.join(instruction_list)} 
}}
'''

with open("instructions.go","w") as f:
    f.write(preface)
    f.write(const_declaration)
    f.write(instruction_map_declaration)