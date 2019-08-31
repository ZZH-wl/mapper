package main

import (
	"fmt"
	"github.com/ZZH-wl/mapper"
	order "github.com/ZZH-wl/mapper/example/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type (
	User struct {
		Name     string
		Age      int
		Id       string `mapper:"_id"`
		AA       string `json:"Score,omitempty"`
		Data     []byte
		Students []Student
		Time     time.Time
	}

	Student struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Score string
	}

	Teacher struct {
		Name  string
		Age   int
		Id    string `mapper:"_id"`
		Level string
	}

	JsonUser struct {
		Name string
		Age  int
		Time mapper.JSONTime
	}
)

func init() {
	mapper.Register(&User{})
	mapper.Register(&Student{})
}

func main() {
	user := &User{}
	userMap := &User{}
	teacher := &Teacher{}
	student := &Student{Name: "test", Age: 10, Id: "testId", Score: "100"}
	valMap := make(map[string]interface{})
	valMap["Name"] = "map"
	valMap["Age"] = 10
	valMap["_id"] = "x1asd"
	valMap["Score"] = 100
	valMap["Data"] = []byte{1, 2, 3, 4}
	valMap["Students"] = []byte{1, 2, 3, 4} //[]Student{*student}
	valMap["Time"] = time.Now()

	mapper.SetEnabledTypeChecking(true)

	mapper.Mapper(student, user)
	mapper.AutoMapper(student, teacher)
	mapper.MapperMap(valMap, userMap)

	fmt.Println("student:", student)
	fmt.Println("user:", user)
	fmt.Println("teacher", teacher)
	fmt.Println("userMap:", userMap)

	jsonUser := &JsonUser{
		Name: "json",
		Age:  1,
		Time: mapper.JSONTime(time.Now()),
	}

	fmt.Println(jsonUser)


	orderObj := &order.Order{
		Id:                   12,
		CreatedOn:            &timestamp.Timestamp{Seconds: time.Now().Unix()},
		ModifiedOn:           &timestamp.Timestamp{Seconds: time.Now().Unix()},
		DeletedOn:            &timestamp.Timestamp{Seconds: time.Now().Unix()},
		PayTime:              &timestamp.Timestamp{Seconds: time.Now().Unix()},
		StoreId:              12,
		UserUid:              "234678fgdxcbvn424467",
		OrderNo:              "78521369542369521525",
		PayNo:                "852374285239984223685",
		LogisticsNo:          "788651245456512",
		OrderStatus:          3.00,
		OriginFee:            100.00,
		PayFee:               120.4521,
		UserMem:              "当时法国卡斯蒂略结合当时法国卡斯蒂略结合体",
		IsStore:              false,
	}

	oldOrder := &order.OldOrder{
		Id:                   12,
		CreatedOn:            time.Now(),
		//ModifiedOn:           time.Now(),
		//DeletedOn:            time.Now(),
		PayTime:              time.Now(),
		StoreId:              12,
		UserUid:              "234678fgdxcbvn424467",
		OrderNo:              "78521369542369521525",
		PayNo:                "852374285239984223685",
		LogisticsNo:          "788651245456512",
		OrderStatus:          3.00,
		OriginFee:            100.00,
		PayFee:               120.3432,
		UserMem:              "当时法国卡斯蒂略结合当时法国卡斯蒂略结合体",
		IsStore:              false,
	}
	order1 := &order.Order{}
	oldOrder1 := &order.OldOrder{}
	order2 := &order.Order{}
	oldOrder2 := &order.OldOrder{}

	mapper.Mapper(oldOrder, order1)
	mapper.Mapper(order1, oldOrder1)

	mapper.AutoMapper(orderObj, oldOrder2)
	mapper.AutoMapper(oldOrder2, order2)
}
