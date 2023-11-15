package main

import (
	"fmt"
	"gin_blog/models"
	"gin_blog/pkg/gredis"
	"gin_blog/pkg/logging"
	"gin_blog/pkg/setting"
	"gin_blog/pkg/util"
	"gin_blog/routers"
)

func main() {
	// endless.DefaultReadTimeOut = setting.ReadTimeOut
	// endless.DefaultWriteTimeOut = setting.WriteTimeOut
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.HttpPort)
	// server := endless.NewServer(endPoint, routers.InitRouters())

	setting.Setup()
	gredis.SetUp()
	util.SetUp()
	logging.Setup()
	models.Setup()
	r := routers.InitRouters()
	//logging.Info(setting.ServerSetting.HttpPort)
	//r.Run(":" + setting.ServerSetting.HttpPort)
	r.Run(fmt.Sprintf(":%v", setting.ServerSetting.HttpPort))

	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Server err: %v", err)
	// }
	//

}

func Lsort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2

	Lsort(arr, l, mid)
	Lsort(arr, mid+1, r)
	//0,2--3,5
	//12345
	//1.12
	//2 .34
	//3 .5
	//0,1--2,2 -- 3,4--5,5
	//
	merge(arr, l, mid, r)

}
func merge(arr []int, l, mid, r int) {
	var aux = make([]int, len(arr))
	copy(aux, arr)
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > r {
			arr[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			arr[k] = aux[i]
			i++
		} else {
			arr[k] = aux[j]
			j++
		}
	}
}
func ShellSort(arr []int) []int {
	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(arr); i++ {
			for j := i; j >= gap; j -= gap {
				if arr[j] < arr[j-gap] {
					arr[j], arr[j-gap] = arr[j-gap], arr[j]
				}

			}

		}
	}
	return arr
}
