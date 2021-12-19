package models

type Database map[string]string

type Store struct {
	Db Database
}
