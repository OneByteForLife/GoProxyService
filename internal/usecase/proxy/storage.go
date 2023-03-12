package proxy

import (
	"GoProxyService/internal/entity"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

// table
const (
	proxy_list = "proxy_list"
)

// cols
const (
	types      = "types"
	ip         = "ip"
	port       = "port"
	speed      = "speed"
	anonlvl    = "anonlvl"
	city       = "city"
	country    = "country"
	last_check = "last_check"
)

type ProxyStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &ProxyStorage{db}
}

func (s *ProxyStorage) SaveData(data []entity.ProxyData) error {
	for _, val := range data {
		ds := goqu.Insert(proxy_list).Cols(types, ip, port, speed, anonlvl, city, country, last_check).Vals(
			goqu.Vals{
				&val.Types[0],
				&val.Data.IP,
				&val.Data.Port,
				&val.Data.Speed,
				&val.Data.AnonLvL,
				&val.Data.Geo.City,
				&val.Data.Geo.Country,
				time.Now().Format("2006-01-02 15:04:05"),
			},
		)

		q, _, _ := ds.ToSQL()

		_, err := s.db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}
