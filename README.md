# json-to-go
Hacky extension to https://github.com/a-h/generate

The 'generate' package does some nice schema->struct conversion. But it can't handle '$ref' elements that point to other files.

This repo alleviates the situation a bit, by inflating the json first (dereferencing every 'ref'), before calling the generator.
So it looks up referenced files, and replaces '$ref' with its content.

## Limitations:
- only works for local files
