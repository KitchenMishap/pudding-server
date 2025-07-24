An http rest server for Bitcoin data gathered by pudding-shed

Example http GET requests:

Ask about the first block
http://127.0.0.1:14699/fulldag-vertex/block/0/attributes.json
http://127.0.0.1:14699/fulldag-vertex/block/0/in.json
http://127.0.0.1:14699/fulldag-vertex/block/0/out.json

Ask about the first block (minidag gives a more concise DAG)
http://127.0.0.1:14699/minidag-vertex/block/0/attributes.json
http://127.0.0.1:14699/minidag-vertex/block/0/in.json
http://127.0.0.1:14699/minidag-vertex/block/0/out.json

Ask about the first transaction
http://127.0.0.1:14699/fulldag-vertex/transaction/0/attributes.json
http://127.0.0.1:14699/fulldag-vertex/transaction/0/in.json
http://127.0.0.1:14699/fulldag-vertex/transaction/0/out.json

Ask about the first txo
http://127.0.0.1:14699/fulldag-vertex/txo/0/attributes.json
http://127.0.0.1:14699/fulldag-vertex/txo/0/in.json
http://127.0.0.1:14699/fulldag-vertex/txo/0/out.json

Ask about the first address
http://127.0.0.1:14699/fulldag-vertex/address/0/attributes.json
http://127.0.0.1:14699/fulldag-vertex/address/0/in.json
http://127.0.0.1:14699/fulldag-vertex/address/0/out.json

Ask about a hash
http://127.0.0.1:14699/lookup/000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f

Ask about an address
http://127.0.0.1:14699/lookup/1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa