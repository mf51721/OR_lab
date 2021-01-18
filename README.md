# Otvoreno Računarstvo - 4. laboratorijska vježba

API koji izlaže bazu podataka programskih jezika.

GET zahtjev na /api/languages/:id/picture vraća sliku resursa s wikipedije, ako postoji.

### Pokretanje

Za pokrenuti API:
```
go run ./main.go
```

Za podignuti DB koristeći docker-compose:
```
cd db
docker-compose up
```

# Metapodaci
Naslov: Baza Podataka Programskih Jezika i Njihovih Svojstava

Otvorena licensa: Creative Commons Zero v1.0 Universal

Naziv autora: Marko Filipović

Verzija skupa podataka: 1.0

Datum objave: 2020-10-26

Jezik u kojemu se nalaze podaci: n/a (Generic)

Format: SQL, CSV, JSON

Tema: Programski Jezici

Ključne riječi: programiranje, računarstvo, inženjerstvo

Atributi koji se nalaze u skupu podataka: 


name | year | wikipedia | imperative | object_oriented | functional | procedural | generic | reflective | creators
--- | --- | --- | --- | --- | --- | --- | --- | --- | ---
Ime jezika | Godina pojavljivanja | Wikipedia stranica | je li jezik imperativan | je li jezik objektno orijentiran | je li jezik funkcijski | je li jezik proceduralan | je li jezik generičan | je li jezik reflektivan | Autori jezika
