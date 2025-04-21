package initialize

import (
	"context"
	"fmt"
	"log"
	"mtb_web/global"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	m := global.Config.Mongo
	// Lấy ATLAS_URI và DB_NAME từ môi trường
	uri := "mongodb+srv://" + m.Name + ":" + m.Pass + "@" + m.Cluster + ".nmpse.mongodb.net/" + m.DbName + "?retryWrites=true&w=majority&appName=" + m.Cluster
	fmt.Println("Kết nối MongoDB Atlas với URI:", uri)
	// Cấu hình kết nối MongoDB
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Lỗi kết nối MongoDB:", err)
	}

	// Kiểm tra kết nối
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB không phản hồi:", err)
	}

	fmt.Println("Kết nối MongoDB Atlas thành công!")

	// Gán database vào biến toàn cục
	global.Mongodb = client.Database(m.DbName)
}
