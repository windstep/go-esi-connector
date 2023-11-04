# Eve Online API (ESI) for golang applications

**CURRENTLY THIS IS A WORK-IN-PROGRESS LIBRARY**

This package is created to simplify Eve Online API usage in Go methods. It is crafted carefully with all community and
CCP guidelines and best practices. This connection layer was originally designed for my fleet waitlist application for Uroborus - 
Sansha ratting community:)

What can you do with it?

- Access to all public (and private if you have token) methods of the ESI API
- Work with pre-defined response objects

How to use this package?

```
package workwithme
import (
	"github.com/windstep/go-esi-connector"
	"net/http"
)

httpClient = &http.Client{}
client := esi.NewClient("https://esi.evetech.net", "Your Application Name", httpClient)
allianceIds, err := client.GetAlliancesIds()
```

Methods in this package are similar to the [Eve Swagger Interface](https://esi.evetech.net). So sometimes you will need to pass a userToken
parameter. Sometimes you will need to work with `*Page` parameter which gives you understanding, that this endpoint could be called
more than ones to obtain full list of data.

About the project structure. I've chosen the flat structure: there are just two kinds of files here. `client.go` is all about how we make a query
and other files are just the representation of high level parts of ESI api - alliances (4 queries), assets (6 queries) etc. But all the
queries are just the methods for the Client object you create.

This package **still** does not work with caching, but it's in my TODO list, so I will add some caching layer between client and request
for you to call methods without any stress-testing ESI servers:)

If you have any questions about the code or suggestions - feel free to create an Issue or make a pull request.

Fly safe!

