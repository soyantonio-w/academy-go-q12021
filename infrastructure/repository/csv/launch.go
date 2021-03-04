package csv

import (
	"encoding/csv"
	"github.com/soyantonio-w/academy-go-q12021/entity"
	"os"
	"strconv"
)

type repository struct {
}

func NewRepository() entity.LaunchRepo {
	return &repository{}
}

func (repo *repository)GetLaunches()([]entity.Launch, error)  {
	csvFile, err := os.Open("data/spacex-launches.csv")

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvFile)

	if _, err := r.Read(); err != nil {
		return nil, err
	}

	var launches []entity.Launch
	for record := readLine(r); record != nil; record = readLine(r){
		launchId, _ := strconv.Atoi(record[0])
		success, _ := strconv.ParseBool(record[6])
		launch := entity.NewLaunch(
			launchId,
			record[1],
			record[2],
			record[3],
			record[4],
			record[5],
			success,
			)
		launches = append(launches, launch)
	}
	return launches, nil
}

func readLine(reader *csv.Reader) (line []string) {
	line, _ = reader.Read()
	return
}