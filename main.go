package main

import (
	"fmt"
	initial "geektrust/intial"
	"geektrust/mappings"
	"log"
	"time"
)

// "bufio"

// "geektrust/course"
// "geektrust/types"
// "log"
// "os"
// "strings"

func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Print("panic recovered", r)
		}
	}()
	initial.LoadEnv()

	redisClient := initial.NewRedisClient()
	if _, err := redisClient.Ping().Result(); err != nil {
		log.Fatal(err)
		time.Sleep(time.Second * 2)
		_, err := redisClient.Ping().Result()
		if err != nil {
			panic(err)
		}

	}

	data, _ := redisClient.Get("goredis").Result()

	fmt.Println(data)

	mappings.CreateUrlMappings()
	mappings.Router.Run(":5000")

	// defer (func() {
	// 	recover()
	// })()

	// reads the arguments passed from command line for execution
	// cliArgs := os.Args[1:]

	// if len(cliArgs) == 0 {
	// 	fmt.Println("Please provide the input file path")
	// 	return
	// }

	// opens the file provides as first argument of command in shell
	// filePath := cliArgs[0]
	// file, err := os.Open("sample_input/input1.txt")

	// if err != nil {
	// 	fmt.Println("Error opening the input file")
	// 	return
	// }

	// defer file.Close()
	// scanner := bufio.NewScanner(file)

	/*
		Initializing maps(courseEmployeeRegMap)for mapping CourseOfferingId with its course data and Registered Employees,
		Similary for (courseRegIdMap) Employee course registration Id with CourseOfferingId
	*/
	// var courses []types.Course
	// courseEmployeeRegMap := make(map[string]types.CourseData)
	// courseRegIdMap := make(map[string]string)

	// for scanner.Scan() {
	// 	args := scanner.Text()
	// 	argList := strings.Fields(args)

	// 	//Executes the commands passed in the arguments
	// 	course.ExecuteCommandsFactory(argList, &courses, courseEmployeeRegMap, courseRegIdMap)
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
