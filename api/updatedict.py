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

def add_definitions(data):
    bigint =  {
      "type": "integer",
      "minimum": 0,
      "x-go-type": {
        "import": {
          "package": "github.com/aeternity/aepp-sdk-go/utils"
        },
        "type": "BigInt"
      }
    }
    data["definitions"]["Amount"] = bigint
    data["definitions"]["Balance"] = bigint
    data["definitions"]["Fee"] = bigint
    data["definitions"]["NameSalt"] = bigint
    return data

def add_references_to_definitions(data):
    swaggerD = DottedDict(data)
    for l in json_objects:
        swaggerD_old = swaggerD[l]
        if 'fee' in l:
            swaggerD[l] = {"$ref": "#/definitions/Fee"}
        elif "balance" in l:
            swaggerD[l] = {"$ref": "#/definitions/Balance"}
        elif "amount" in l:
            swaggerD[l] = {"$ref": "#/definitions/Amount"}
        elif "channel_reserve" in l:
            swaggerD[l] = {"$ref": "#/definitions/Amount"}
        elif "name_salt" in l:
            swaggerD[l] = {"$ref": "#/definitions/NameSalt"}
        
        # I want to see what was changed.
        if swaggerD[l] != swaggerD_old:
            print(l, swaggerD[l])
    return swaggerD.to_python()

with open('swaggerOrig.json') as f:
    swagger = json.load(f)
traverse(swagger, [])


swagger_n = add_definitions(swagger)
swagger_n = add_references_to_definitions(swagger_n)
swagger_n = no_implicit_int64(swagger_n)


with open('swaggerD.json', 'w') as fnew:
    json.dump(swagger_n, fnew, indent=2)