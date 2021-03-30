package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/soyantonio-w/academy-go-q12021/entity"
)

type repository struct {
	folder, filename string
}

func NewRepository() entity.LaunchRepo {
	return &repository{
		folder:   "data",
		filename: "spacex-launches.csv",
	}
}

func (repo *repository) Get(id entity.LaunchId) (entity.Launch, error) {
	launches, err := repo.GetLaunches()
	if err != nil {
		return entity.Launch{}, fmt.Errorf("non available launches")
	}

	for _, launch := range launches {
		if launch.LaunchId == id {
			return launch, nil
		}
	}

	return entity.Launch{}, fmt.Errorf("non found launch with id %d", id)
}

func (repo *repository) GetLaunches() ([]entity.Launch, error) {
	csvFile, err := os.Open(repo.getPath())

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(csvFile)

	if _, err := r.Read(); err != nil {
		return nil, err
	}

	var launches []entity.Launch
	for record := readLine(r); record != nil; record = readLine(r) {
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

func (repo *repository) SyncAll(launches []entity.Launch) error {
	filename := fmt.Sprintf("spacex-launches-%d.csv", time.Now().Unix())
	file, err := os.Create(repo.buildPath(filename))
	if err != nil {
		return fmt.Errorf("could not create file")
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, launch := range launches {
		record := []string{
			strconv.Itoa(int(launch.LaunchId)),
			string(launch.LaunchDetails),
			string(launch.LaunchDate),
			string(launch.MissionName),
			string(launch.VideoLink),
			string(launch.RocketName),
			strconv.FormatBool(bool(launch.LaunchSuccess)),
		}
		if err = w.Write(record); err != nil {
			return fmt.Errorf("error writing record to file %s", filename)
		}
	}

	repo.filename = filename
	return nil
}

func (repo *repository) getPath() string {
	return repo.buildPath(repo.filename)
}

func (repo *repository) buildPath(filename string) string {
	return filepath.Join(repo.folder, filename)
}

func readLine(reader *csv.Reader) (line []string) {
	line, _ = reader.Read()
	return
}
