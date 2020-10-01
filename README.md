# go-bsvrates
> Get the current exchange rate of BSV and other helpful currency conversions

[![Release](https://img.shields.io/github/release-pre/tonicpow/go-bsvrates.svg?logo=github&style=flat&v=2)](https://github.com/tonicpow/go-bsvrates/releases)
[![Build Status](https://travis-ci.com/tonicpow/go-bsvrates.svg?branch=master&v=2)](https://travis-ci.com/tonicpow/go-bsvrates)
[![Report](https://goreportcard.com/badge/github.com/tonicpow/go-bsvrates?style=flat&v=2)](https://goreportcard.com/report/github.com/tonicpow/go-bsvrates)
[![codecov](https://codecov.io/gh/tonicpow/go-bsvrates/branch/master/graph/badge.svg?v=2)](https://codecov.io/gh/tonicpow/go-bsvrates)
[![Go](https://img.shields.io/github/go-mod/go-version/tonicpow/go-bsvrates?v=2)](https://golang.org/)

<br/>

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

<br/>

## Installation

**go-bsvrates** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```shell script
go get -u github.com/tonicpow/go-bsvrates
```

<br/>

## Documentation
View the generated [documentation](https://pkg.go.dev/github.com/tonicpow/go-bsvrates)

[![GoDoc](https://godoc.org/github.com/tonicpow/go-bsvrates?status.svg&style=flat)](https://pkg.go.dev/github.com/tonicpow/go-bsvrates)

### Features
- [Client](client.go) is completely configurable
- Using default [heimdall http client](https://github.com/gojek/heimdall) with exponential backoff & more
- Use your own HTTP client
- Helpful currency conversion and formatting methods:
    - [ConvertFloatToIntBSV()](currency.go)
    - [ConvertIntToFloatUSD()](currency.go)
    - [ConvertPriceToSatoshis()](currency.go)
    - [ConvertSatsToBSV()](currency.go)
    - [FormatCentsToDollars()](currency.go)
    - [TransformCurrencyToInt()](currency.go)
    - [TransformIntToCurrency()](currency.go)
- Supported Currencies:
    - USD
- Supported Providers:
    - [Coin Paprika](https://api.coinpaprika.com/)
    - [What's On Chain](https://developers.whatsonchain.com/)
    - [Preev](https://preev.pro/api/)

<details>
<summary><strong><code>Library Deployment</code></strong></summary>
<br/>

[goreleaser](https://github.com/goreleaser/goreleaser) for easy binary or library deployment to Github and can be installed via: `brew install goreleaser`.

The [.goreleaser.yml](.goreleaser.yml) file is used to configure [goreleaser](https://github.com/goreleaser/goreleaser).

Use `make release-snap` to create a snapshot version of the release, and finally `make release` to ship to production.
</details>

<details>
<summary><strong><code>Makefile Commands</code></strong></summary>
<br/>

View all `makefile` commands
```shell script
make help
```

List of all current commands:
```text
clean                  Remove previous builds and any test cache data
clean-mods             Remove all the Go mod cache
coverage               Shows the test coverage
godocs                 Sync the latest tag with GoDocs
help                   Show this help message
install                Install the application
install-go             Install the application (Using Native Go)
lint                   Run the Go lint application
release                Full production release (creates release in Github)
release                Runs common.release then runs godocs
release-snap           Test the full release (build binaries)
release-test           Full production test release (everything except deploy)
replace-version        Replaces the version in HTML/JS (pre-deploy)
run-examples           Runs the basic example
tag                    Generate a new tag and push (tag version=0.0.0)
tag-remove             Remove a tag if found (tag-remove version=0.0.0)
tag-update             Update an existing tag to current commit (tag-update version=0.0.0)
test                   Runs vet, lint and ALL tests
test-short             Runs vet, lint and tests (excludes integration tests)
test-travis            Runs all tests via Travis (also exports coverage)
test-travis-short      Runs unit tests via Travis (also exports coverage)
uninstall              Uninstall the application (and remove files)
vet                    Run the Go vet application
```
</details>

<br/>

## Examples & Tests
All unit tests and [examples](examples) run via [Travis CI](https://travis-ci.org/tonicpow/go-bsvrates) and uses [Go version 1.15.x](https://golang.org/doc/go1.15). View the [deployment configuration file](.travis.yml).

Run all tests (including integration tests)
```shell script
make test
```

Run tests (excluding integration tests)
```shell script
make test-short
```

<br/>

## Benchmarks
Run the Go [benchmarks](client.go):
```shell script
make bench
```

<br/>

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

<br/>

## Usage
View the [examples](examples)

Basic exchange rate implementation:
```go
package main

import (
	"log"

	"github.com/tonicpow/go-bsvrates"
)

func main() {

	// Create a new client (all default providers)
	client := bsvrates.NewClient(nil, nil)
    
	// Get rates
	rate, provider, _ := client.GetRate(bsvrates.CurrencyDollars)
	log.Printf("found rate: %v %s from provider: %s", rate, bsvrates.CurrencyToName(bsvrates.CurrencyDollars), provider.Name())
}
``` 

Basic price conversion implementation:
```go
package main

import (
	"log"

	"github.com/tonicpow/go-bsvrates"
)

func main() {

	// Create a new client (all default providers)
	client := bsvrates.NewClient(nil, nil)
    
	// Get a conversion from $ to Sats
	satoshis, provider, _ := client.GetConversion(bsvrates.CurrencyDollars, 0.01)
	log.Printf("0.01 USD = satoshis: %d from provider: %s", satoshis, provider.Name())
}
```
 
<br/>

## Maintainers
| [<img src="https://github.com/mrz1836.png" height="50" alt="MrZ" />](https://github.com/mrz1836) |
|:---:|
| [MrZ](https://github.com/mrz1836) |
              
<br/>

## Contributing
View the [contributing guidelines](CONTRIBUTING.md) and please follow the [code of conduct](CODE_OF_CONDUCT.md).

### How can I help?
All kinds of contributions are welcome :raised_hands:! 
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:. 
You can also support this project by [becoming a sponsor on GitHub](https://github.com/sponsors/mrz1836) :clap: 
or by making a [**bitcoin donation**](https://tonicpow.com/?af=go-bsvrates) to ensure this journey continues indefinitely! :rocket:


### Credits

[Coin Paprika](https://tncpw.co/7c2cae76), [What's On Chain](https://tncpw.co/638d8e8a) and [Preev](https://tncpw.co/d19f43a3) for their hard work on their public API

[Jad](https://github.com/jadwahab) for his contributions to the package!

<br/>

## License

![License](https://img.shields.io/github/license/tonicpow/go-bsvrates.svg?style=flat&v=2)