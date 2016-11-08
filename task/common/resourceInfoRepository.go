package common

import (
	"database/sql"
	"encoding/json"
)

type ResourceInfoRepository struct {
	CustomStreamsFunc CustomStreams
}

type CustomStreams func(s []byte) map[string]Stream

func (r *ResourceInfoRepository) AddOne(entity ResourceInfo) {
	oldEntity := r.EntityByUrl(entity.Url)
	streams, err := json.Marshal(entity.Streams)
	checkErr(err)

	if oldEntity.Id == 0 {
		// do insert
		stmt, err := conn.Prepare("INSERT resource_info SET site=?,title=?,url=?,streams=?")
		checkErr(err)

		res, err := stmt.Exec(entity.Site, entity.Title, entity.Url, string(streams))
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)
		entity.Id = id
	} else {
		// do update
		entity.Id = oldEntity.Id
		stmt, err := conn.Prepare("UPDATE resource_info SET site=?,title=?,streams=? where id=?")
		checkErr(err)

		_, err = stmt.Exec(entity.Site, entity.Title, streams, entity.Id)
		checkErr(err)
	}
}

func (r *ResourceInfoRepository) SelectInfoFromDb(where interface{}) ResourceInfo {
	var id int64
	var site string
	var title string
	var url string
	var streams string
	var err error

	switch where.(type) {
	case int64:
		err = conn.QueryRow("SELECT * FROM resource_info where id=?", where).Scan(&id, &site, &title, &url, &streams)
	case string:
		err = conn.QueryRow("SELECT * FROM resource_info where url=?", where).Scan(&id, &site, &title, &url, &streams)
	default:
		return ResourceInfo{}
	}

	if err == sql.ErrNoRows {
		return ResourceInfo{}
	}
	checkErr(err)

	streamsMap := r.CustomStreamsFunc([]byte(streams))

	resInfo := ResourceInfo{
		Id:      id,
		Site:    site,
		Title:   title,
		Url:     url,
		Streams: streamsMap,
	}

	return resInfo
}

func (r *ResourceInfoRepository) Entity(id int64) ResourceInfo {
	return r.SelectInfoFromDb(id)
}

func (r *ResourceInfoRepository) EntityByUrl(u string) ResourceInfo {
	return r.SelectInfoFromDb(u)
}
