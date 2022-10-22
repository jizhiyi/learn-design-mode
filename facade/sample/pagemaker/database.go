package pagemaker

import "learn/adapter/a2"

func getProperties(dbname string) (a2.FileIO, error) {
	fileName := dbname + ".txt"
	fileProp := a2.NewFileProperties()
	err := fileProp.ReadFromFile(fileName)
	if err != nil {
		return nil, err
	}
	return fileProp, nil
}
