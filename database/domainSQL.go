package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func CreateTable(db *sql.DB) {

	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS domain (host STRING PRIMARY KEY, sslGrade STRING, previousGrade STRING, lastSearch TIMESTAMPTZ)"); err != nil {
		log.Fatal(err)
	}

}
func InsertDomain(db *sql.DB, host string, sslgrade string, previous string) {
	parameters := "'" + host + "','" + sslgrade + "','" + previous + "', "
	if _, err := db.Exec(
		"INSERT INTO domain (host, sslGrade, previousGrade, lastSearch) VALUES (" + parameters + " NOW())"); err != nil {
		log.Fatal(err)
	}

}

func UpdateDomain(db *sql.DB, host string, sslgrade string, previous string) {
	parameters := "'" + sslgrade + "','" + previous + "', "
	if _, err := db.Exec(
		"UPDATE domain SET (sslGrade, previousGrade, lastSearch) = ("+parameters+" NOW()) WHERE host = $1", host); err != nil {
		log.Fatal(err)
	}

}
func GetDomains(db *sql.DB) []DomainDB {
	var domains []DomainDB
	rows, err := db.Query("SELECT * FROM domain ORDER BY lastSearch DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var host, sslGrade, previousGrade string
		var lastSearch time.Time
		if err := rows.Scan(&host, &sslGrade, &previousGrade, &lastSearch); err != nil {
			log.Fatal(err)
		}
		domain := DomainDB{host, sslGrade, previousGrade, lastSearch}
		domains = append(domains, domain)
	}
	return domains
}

func SearchDomain(db *sql.DB, host string) DomainDB {
	//var domain DomainDB
	rows, err := db.Query("SELECT * FROM domain WHERE host = $1", host)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var host, sslGrade, previousGrade string
		var lastSearch time.Time
		if err := rows.Scan(&host, &sslGrade, &previousGrade, &lastSearch); err != nil {
			log.Fatal(err)
		}
		return DomainDB{Host: host, SslGrade: sslGrade, PreviousSSL: previousGrade, LastTime: lastSearch}
		//fmt.Println(domain)
		//fmt.Println(domain.Host)
	}
	return DomainDB{}

}
