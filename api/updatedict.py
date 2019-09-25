from pprint import pprint
from dotted.collection import DottedDict
import json

json_objects = []
json_leaves = []
def traverse(data, parents):
    for key in data:
        # If we're in an "example": {"fee": 5} object, don't traverse into it.
        if key == "example":
            continue 
        if not isinstance(data[key], dict):
            json_leaves.append(".".join(parents + [key]))
            continue

        json_objects.append(".".join(parents + [key]))
        parents.append(key)
        traverse(data[key], parents)
    
    try:
        parents.pop()
    except IndexError:
        # We have reached the JSON's root.
        pass

def no_implicit_int64(data):
    swaggerD = DottedDict(data)
    for l in json_objects:
        n = swaggerD[l]
        if n.get("type") == "integer" and n.get("format") is None:
            n.format = "uint64"
            print(l, n)
    
    return swaggerD.to_python()

def add_uint_bigint(data):
    bigint =  {
      "type": "integer",
      "minimum": 0,
      "x-go-type": {
        "import": {
          "package": "github.com/aeternity/aepp-sdk-go/v5/utils"
        },
        "type": "BigInt"
      }
    }
    data["definitions"]["UInt"] = bigint
    return data

with open('rc2.json') as f:
    swagger = json.load(f)
traverse(swagger, [])


swagger_n = add_uint_bigint(swagger)
swagger_n = no_implicit_int64(swagger_n)


with open('swaggerD.json', 'w') as fnew:
    json.dump(swagger_n, fnew, indent=2)