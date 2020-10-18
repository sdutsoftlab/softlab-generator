package gen

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var (
	log = logrus.WithFields(logrus.Fields{
		`MOD`: `gen`,
	})
)

func ListenHTTPServer(port int) {
	log.Info("打开内置服务器")

	http.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir(conf.Dist+"/assets/"))))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(conf.Dist))))

	//log.Info("内置web服务器开启成功，监听端口 :%d...", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
