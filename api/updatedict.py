import json
import argparse

def no_implicit_int64(el, p = []):
    if isinstance(el, list):
        for i, d in enumerate(el):
            no_implicit_int64(d, p + [str(i)])
    if isinstance(el, dict):
        if el.get('type') == 'integer' and el.get('format') is None:
            if el.get('minimum', 0) >= 0:
                el['format'] = 'uint64'
                print('added uint64 in', '.'.join(p), el)
            el.pop('maximum', None)
        for k, v in el.items():
            no_implicit_int64(v, p + [k])

replaces = [
    ['"$ref": "#/definitions/UInt64"', '"type":"integer", "format":"uint64"'],
    ['"$ref": "#/definitions/UInt32"', '"type":"integer", "format":"uint32"'],
    ['"$ref": "#/definitions/UInt16"', '"type":"integer", "format":"uint16"'],

    ['"$ref": "#/definitions/EncodedPubkey"', '"type":"string"'],
    ['"$ref": "#/definitions/EncodedHash"', '"type":"string"'],
    ['"$ref": "#/definitions/EncodedValue"', '"type":"string"'],
    ['"$ref": "#/definitions/EncodedByteArray"', '"type":"string"'],

    ['"#/definitions/TxBlockHeight"', '"#/definitions/UInt"'],
    ['OracleResponseTxJSON', 'OracleRespondTxJSON'],
]

parser = argparse.ArgumentParser(description='Modify a node.json')
parser.add_argument('infile', type=open)
parser.add_argument('outfile', type=argparse.FileType('w'))
args = parser.parse_args()

api = args.infile.read()

for replace in replaces:
    api = api.replace(*replace)

parsed = json.loads(api)

no_implicit_int64(parsed)

parsed['definitions']['UInt'] = {
  'type': 'integer',
  'minimum': 0,
  'x-go-type': {
    'import': {
      'package': 'github.com/aeternity/aepp-sdk-go/v9/utils'
    },
    'type': 'BigInt'
  }
}

json.dump(parsed, args.outfile, indent=2)
