/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 10:56
 */
package main

import (
	"flag"
	"fmt"
	"gochat/connect"
	"gochat/job"
	"gochat/logic"
)

func main() {
	var module string
	flag.StringVar(&module, "module", "", "assign run module")
	flag.Parse()
	fmt.Println(fmt.Sprintf("start run %s module", module))
	switch module {
	case "connect":
		connect.New().Run()
	case "logic":
		logic.New().Run()
	case "job":
		job.New().Run()
	default:
		fmt.Println("exiting,module param error!")
		return
	}
	fmt.Println(fmt.Sprintf("run %s module done!", module))
}