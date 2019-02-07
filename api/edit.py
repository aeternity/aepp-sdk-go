import json

def traverse(data):
    try:
        for key in data:
            # If we're in an "example": {"fee": 5} object, don't traverse into it.
            if key == "example":
                continue
            v = data[key]
            data[key] = replace_fee(key, v)
            traverse(v)
    except TypeError:
        pass
    return data

def replace_fee(key, value):
    if "fee" in key:
        value.pop("type", None)
        value.pop("format", None)
        value["$ref"] = "#/definitions/Fee"
        return value
    return value

with open('swagger.json') as f:
    api = json.load(f)

api_rewritten = traverse(api)

with open('swagger-rewrite.json', 'w') as fnew:
    json.dump(api_rewritten, fnew, indent=2)