# Fietsverhuur van der Binckes
Om het beroepsproduct gebruiksklaar te maken:
1. Plaats de map van-der-binckes in uw `GOPATH`
2. ``go get github.com/go-sql-driver/mysql``
3. ``go build van-der-binckes`` 
4. Maak in uw host (localhost, WAMP,XAMPP, LAMP of een online host) een mysql database genaamd `vanderbinckes`
5. Importeer de sql uit de assets map
6. ``go run van-der-binckes.go``
