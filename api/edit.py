"""
1. Manually replace all int64s with uint64s in swagger.json except for time (because go's time.Unix() accepts int64, not uint64)
2. Run this to add Fee/Balance/Amount BigInt in definitions and replace where necessary

CAVEATS
/generations/height/{height} cannot be targeted, have to replace by hand
OffChain* allOf is a list, this script cannot target amounts and fees in there
some heights are int64/uint64
"""
import argparse
import json

def traverse(data):
    try:
        for key in data:
            # If we're in an "example": {"fee": 5} object, don't traverse into it.
            if key == "example":
                continue

            v = data[key]
            if "fee" in key:
                v.pop("type", None)
                v.pop("format", None)
                v.pop("minimum", None)
                v["$ref"] = "#/definitions/Fee"
            elif "balance" == key:
                v.pop("type", None)
                v.pop("format", None)
                v.pop("minimum", None)
                v["$ref"] = "#/definitions/Balance"
            elif "amount" in key:
                v.pop("type", None)
                v.pop("format", None)
                v.pop("minimum", None)
                v["$ref"] = "#/definitions/Amount"
            elif "channel_reserve" == key:
                v.pop("type", None)
                v.pop("format", None)
                v.pop("minimum", None)
                v["$ref"] = "#/definitions/Amount"
            elif "block_height" == key:
                v["format"] = "uint64"
            elif "name_salt" == key:
                v.pop("type", None)
                v.pop("format", None)
                v["$ref"] = "#/definitions/NameSalt"
                
            traverse(v)
    except TypeError:
        pass
    return data

def add_definitions(data):
    data["definitions"]["Amount"] = {
      "type": "integer",
      "minimum": 0,
      "x-go-type": {
        "import": {
          "package": "github.com/aeternity/aepp-sdk-go/utils"
        },
        "type": "BigInt"
      }
    }

    data["definitions"]["Balance"] = {
      "type": "integer",
      "minimum": 0,
      "x-go-type": {
        "import": {
          "package": "github.com/aeternity/aepp-sdk-go/utils"
        },
        "type": "BigInt"
      }
    }

    data["definitions"]["Fee"] = {
      "type": "integer",
      "minimum": 0,
      "x-go-type": {
        "import": {
          "package": "github.com/aeternity/aepp-sdk-go/utils"
        },
        "type": "BigInt"
      }
    }

    data["definitions"]["NameSalt"] = {
      "type": "integer",
      "minimum": 0,
      "x-go-type": {
        "import": {
          "package": "github.com/aeternity/aepp-sdk-go/utils"
        },
        "type": "BigInt"
      }
    }

    return data



parser = argparse.ArgumentParser(description='Automate some (not all) swagger.json type conversions')
parser.add_argument('input', type=str, help='the original swagger.json')
parser.add_argument('output', type=str, help='filename to save modified swagger.json under')
args = parser.parse_args()

with open(args.input) as f:
    api = json.load(f)

api_amount_balance_fee = add_definitions(api)
api_rewritten = traverse(api_amount_balance_fee)

with open(args.output, 'w') as fnew:
    json.dump(api_rewritten, fnew, indent=2)
