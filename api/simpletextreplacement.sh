sed -e 's/"$ref": "#\/definitions\/UInt64"/"type":"integer", "format":"uint64"/g' $1 |
sed -e 's/"$ref": "#\/definitions\/UInt32"/"type":"integer", "format":"uint32"/g' |
sed -e 's/"$ref": "#\/definitions\/UInt16"/"type":"integer", "format":"uint16"/g' |

sed -e 's/"$ref": "#\/definitions\/EncodedPubkey"/"type":"string"/g' |
sed -e 's/"$ref": "#\/definitions\/EncodedHash"/"type":"string"/g' |
sed -e 's/"$ref": "#\/definitions\/EncodedValue"/"type":"string"/g' |
sed -e 's/"$ref": "#\/definitions\/EncodedByteArray"/"type":"string"/g'
