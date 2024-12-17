package mock

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionTemplate *mongo.Collection = configs.GetCOllection(configs.Client, "templates")
func JenisUserSeeder() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newJenisUser := []interface{}{
		models.TemplateRequest{
			JenisUser : "mahasiswa",
			Template : []primitive.ObjectID{},
			CreatedAt : time.Now(),
			UpdatedAt : time.Now(),
		},
		models.TemplateRequest{
			JenisUser : "dosen",
			Template : []primitive.ObjectID{},
			CreatedAt : time.Now(),
			UpdatedAt : time.Now(),
		},
		models.TemplateRequest{
			JenisUser: "tendik",
			Template: []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			JenisUser : "kps",
			Template : []primitive.ObjectID{},
			CreatedAt : time.Now(),
			UpdatedAt : time.Now(),
		},
		models.TemplateRequest{
			JenisUser: "dekanat",
			Template: []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			JenisUser: "ketua_unit",
			Template: []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			JenisUser : "pimpinan_univ",
			Template : []primitive.ObjectID{},
			CreatedAt : time.Now(),
			UpdatedAt : time.Now(),
		},
	}

	_, err := collectionTemplate.InsertMany(ctx, newJenisUser)
	if err != nil {
		panic(err)
	}

	fmt.Println("Jenis User Seeder Created")
}


