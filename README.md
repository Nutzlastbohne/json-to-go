# json-to-go
Hacky extension to https://github.com/a-h/generate

The 'generate' package does some nice schema->struct conversion. But it can't handle '$ref' elements that point to other files.

This repo alleviates the situation a bit, by inflating the json first, before calling the generator.
So it looks up referenced files, and replaces '$ref' with its content.

## Limitations:
- only works for local files
- can't handle references to specific nodes in other files

So, this is handled
```json
"$ref": "metadata.json"
```
this isn't
```json
"$ref": "metadata.json#/definitions/address"
```

## TODO
- make it work with references to specific nodes in other files
