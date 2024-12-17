package mock

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var collectionModule *mongo.Collection = configs.GetCOllection(configs.Client, "module")

func ModuleSeeder() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newModule := []interface{}{
		models.ModuleRequest{
			NAMA_MODULE: "Satu data",
			KETERANGAN: "Satu data",
			URUTAN: 1,
			ICON:       "",
			STATUS:    true,
			CREATED_AT: time.Now().Format("2006-01-02 15:04:05"),
			UPDATED_AT: time.Now().Format("2006-01-02 15:04:05"),
		},
		models.ModuleRequest{
			NAMA_MODULE: "Mahasiswa",
			KETERANGAN: "Mahasiswa",
			URUTAN: 2,
			ICON:       "",
			STATUS:    true,
			CREATED_AT: time.Now().Format("2006-01-02 15:04:05"),
			UPDATED_AT: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	_, err := collectionModule.InsertMany(ctx, newModule)

	if err != nil {
		panic(err)
	}

	fmt.Print("Module Seeder Created")
}