# Registry Contract [![Build Status](https://travis-ci.org/copyright-project/registry-contract.svg?branch=master)](https://travis-ci.org/copyright-project/registry-contract)
> Registry smart contract

## API
### `registerMedia(pHash, imageURL, postedAt, copyrights, binaryHash string)`
Method registers provided media information in the registry.

|Argument Name|Argument Type|Description|
|-------------|-------------|-----------|
|pHash|String|Media's perceptual hash (64 bit)|
|imageUrl|String|Media URL|
|postedAt|String|Timestamp when media was created/posted|
|copyrights|String|Media's copyright attribution|
|binaryHash|String|Binary hash of media's file|

Media's phash, depending on the used algorithm, might produce collisions. Therefore, if media with the same phash is registered, at least binary hash must be different, otherwise it will result in exception.

### `getMedia(pHash string) []string`
Returns the list of registered medias corresponding to a given phash. 
E.g. 
```
[
    "https://some-url-1,123456789,by me,eQYhYzRyWJjPjzpfRFEgmotaFetHsbZRjxAwnwekrBEmfdzdcEkXBAkjQZLCtTMt", 
    "https://some-url-2,123456789,by me,TCoaNatyyiNKAReKJyiXJrscctNswYNsGRussVmaozFZBsbOJiFQGZsnwTKSmVoi"
]
```