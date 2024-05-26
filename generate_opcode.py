import pandas as pd
df = pd.read_csv('opcodes.csv')

func_template = '''
// Op{} executes the {} instruction ({}). {}
func (c *CPU) Op{}() error {{
    panic("not implemented")
    return nil
}}
'''

def display_flags(letters):
    if letters == 'All':
        return letters
    else:
        return ', '.join(list(letters))

complete_file = []
for _, row in df.iterrows():
    opcode, name, flags = [(el if type(el) == str else '').strip() for el in row]
    flags_string = f"Flags: {display_flags(flags)}" if flags != '' else ""
    complete_file.append(func_template.format(opcode,opcode,name,flags_string,opcode))

with open("opcode.go.template","w") as f:
    f.write('\n'.join(complete_file))
