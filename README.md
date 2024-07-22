# GoFaker

GoFaker is library for generating fake data in Go. It provides an API for generating the following fake data:

- organization:
  - `INN` - [INN](https://ru.wikipedia.org/wiki/Идентификационный_номер_налогоплательщика) of private person
  - `INNLE` - [INN](https://ru.wikipedia.org/wiki/Идентификационный_номер_налогоплательщика) of legal entity
  - `OGRN` - [OGRN](https://ru.wikipedia.org/wiki/Основной_государственный_регистрационный_номер) of legal entity
  - `OGRNIP` - [OGRN](https://ru.wikipedia.org/wiki/Основной_государственный_регистрационный_номер) of individual entrepreneur
  - `SNILS` - [SNILS](https://ru.wikipedia.org/wiki/Страховой_номер_индивидуального_лицевого_счёта)

## Installation

To install GoFaker, simply run the following command in your terminal:

```shell
go get -u github.com/Gwinkamp/gofaker
```

## Usage

Here's an example of how to use GoFaker to generate fake data:

```go
package main

import (
	"fmt"

	"github.com/Gwinkamp/gofaker/org"
)

func main() {
	faker := org.NewFaker()

	inn := faker.INN()
	innLE := faker.INNLE()
	ogrn := faker.OGRN()
	ogrnIP := faker.OGRNIP()
	snils := faker.SNILS()

	fmt.Println(inn.String())
	fmt.Println(innLE.String())
	fmt.Println(ogrn.String())
	fmt.Println(ogrnIP.String())
	fmt.Print(snils.String())
}
```

## Contributing

I welcome contributions to GoFaker! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
