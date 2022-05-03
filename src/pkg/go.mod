//module gorm
module github.com/dna/gorm

go 1.17

require (
	github.com/gorilla/mux v1.8.0
	github.com/segmentio/kafka-go v0.4.31
	gorm.io/driver/postgres v1.3.5
	gorm.io/gorm v1.23.4
)

require (
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/klauspost/compress v1.14.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/text v0.3.7 // indirect
)

//replace github.com/fiserv/dna/my_srv1 => ./my_srv1
//I've got one module with mulitple packages
//replace github.com/fiserv/dna/gorm/pkg/handlers => ./pkg/handlers
//replace github.com/fiserv/dna/gorm/pkg/db => ./pkg/db
replace github.com/dna/gorm/pkg/handlers => ./pkg/handlers
replace github.com/dna/gorm/pkg/db => ./pkg/db
replace github.com/dna/gorm/pkg/models => ./pkg/models

require (
	github.com/dna/gorm/pkg/handlers v1.0.0
	github.com/dna/gorm/pkg/db v1.0.0
	github.com/dna/gorm/pkg/models v1.0.0
)



// "gorm/pkg/handlers"
// 	"gorm/pkg/db"
