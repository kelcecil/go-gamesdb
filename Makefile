all : generate

generate :
	go install github.com/kelcecil/go-gamesdb/generategamesdb
	go generate

clean :
	rm get_*.go	

