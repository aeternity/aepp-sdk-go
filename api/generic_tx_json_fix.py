import argparse
import glob
import os
import re

parser = argparse.ArgumentParser(description='Fix generic txs not deserializing properly into their specific Tx types. Edit generic_tx.go to pick up the changes made by this script')
parser.add_argument('directory', type=str, help='the directory to the generated/models/*_tx_json.go files')
args = parser.parse_args()

os.chdir(args.directory)
files_to_edit = glob.glob('*_tx_json.go')

p = re.compile("return \".*TxJSON")
for filename in files_to_edit:
    with open(filename, 'r') as f:
        t = f.read()
    result = p.search(t)
    if result:
        found = result.group(0)
        t_modified = t.replace(found, found[0:-4])
        with open(filename, 'w') as f:
            f.write(t_modified)
            print(filename, ": replaced {} with {}".format(found, found[0:-4]))
