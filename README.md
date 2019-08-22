# Registry Contract [![Build Status](https://travis-ci.org/copyright-project/registry-contract.svg?branch=master)](https://travis-ci.org/copyright-project/registry-contract)
> Registry smart contract

## API
### `registerMedia(mediaID, metadata string)`
Method that registers a metadata (JSON object) with the given media id. <br />
Metadata object structure:
```ts
[media_id: string]: {
    imageUrl: string
    postUrl: string
    postedAt: string
    ownerId: string
    hash: string
}
```

### `areRegistered(ids string) string`
Method that checks whether the list of passed media ids are already registered. <br />
Due to limitations of smart contracts API, arguments and response are in the following string format. <br />
E.g.
```js
areRegistered("id1,id2,id3") 
// "101"
```
`1` means that id corresponding to a position of `1` is registered,`0` - id is not registered.