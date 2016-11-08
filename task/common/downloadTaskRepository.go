package common

type DownloadTaskRepository struct {
}

func (r *DownloadTaskRepository) Add(info ResourceInfo, sIndex string) DownloadTask {
	// do insert
	stmt, err := conn.Prepare("INSERT download_task SET resource_info_id=?,stream_index=?,status=?")
	checkErr(err)

	res, err := stmt.Exec(info.Id, sIndex, sReady)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return DownloadTask{Id: id, Info: &info, InfoId: info.Id, SIndex: sIndex, Status: sReady}
}
