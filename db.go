package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getCaches() (DBCaches, error) {
	var caches DBCaches

	db, err := sqlx.Connect("postgres", "user=audiocache password=audiocache dbname=audiocache sslmode=disable")
	if err != nil {
		return caches, err
	}

	err = db.Select(&caches, "SELECT * FROM caches ORDER BY id ASC")
	if err != nil {
		return caches, err
	}

	err = db.Close()
	return caches, err
}

func getCache(id uint64) (DBCache, error) {
	var cache DBCache

	db, err := sqlx.Connect("postgres", "user=audiocache password=audiocache dbname=audiocache sslmode=disable")
	if err != nil {
		return cache, err
	}

	err = db.Get(&cache, "SELECT * FROM caches WHERE id=$1", id)
	if err != nil {
		return cache, err
	}

	err = db.Close()
	return cache, err
}

func postCache(dbcache DBCache) error {
	db, err := sqlx.Connect("postgres", "user=audiocache password=audiocache dbname=audiocache sslmode=disable")
	if err != nil {
		return err
	}

	_, err = db.NamedExec("INSERT INTO caches (latitude, longitude, created, path) VALUES (:latitude, :longitude, :created, :path)", dbcache)
	if err != nil {
		return err
	}

	err = db.Close()
	return err
}
