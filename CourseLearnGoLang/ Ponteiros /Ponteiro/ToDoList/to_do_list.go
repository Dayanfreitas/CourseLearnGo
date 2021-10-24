package main

import (
	"fmt"
	"reflect"
)

const first = iota + 1

var fake_id int

type Task struct {
	Id    int
	Title string
	Next  *Task
}

type Jobs struct {
	List []*Task
}

func main() {
	// var task *Task
	var op string
	var jobs Jobs

	jobs = Jobs{
		List: make([]*Task, 0),
	}

	for {
		fmt.Println()
		task := createTask()
		jobs.addTask(task)

		fmt.Printf("Para [C]cotinuar [s]sair _> ")
		fmt.Scanln(&op)
		if op == "S" || op == "s" {
			break
		}
	}

	fmt.Println(" ")
	fmt.Println("Jobs _>", jobs)

	for _, v := range jobs.List {
		// if (*v).Next != nil {
		// 	fmt.Println("Task: ", v.Title, "->", (*v).Next.Title)
		// } else {
		// 	fmt.Println("Task: ", v.Title)
		// }

		fmt.Println("#", v.Id, "Task: ", v.Title)
	}
	jobs.removeTask(2)
	// for _, v := range jobs.List {
	// if (*v).Next != nil {
	// 	fmt.Println("Task: ", v.Title, "->", (*v).Next.Title)
	// } else {
	// 	fmt.Println("Task: ", v.Title)
	// }

	// fmt.Println("#", v.Id, "Task: ", v.Title)
	// }

}

func (t *Task) prev() {

}

func (j Jobs) prev() *Task {

	return j.List[len(j.List)-1]
}

func (j *Jobs) addTask(t *Task) {
	if len(j.List) >= first {
		j.prev().Next = t
	}

	j.List = append((j.List), t)
}

func (j Jobs) removeTask(id int) {
	idx := Find(j.List, func(value interface{}) bool {
		return value.(*Task).Id == id
	})

	fmt.Println("INDEX", idx)

	fmt.Println("NEXT", j.List[idx])
	j.List[idx-1].Next = j.List[idx].Next
	j.List[idx].Next = nil
	// j.List[idx].Next.prev() = // j.List[idx].Next.next()

}

func createTask() *Task {
	fake_id++
	var title string

	fmt.Println("Create Task")
	fmt.Println("Title:")
	fmt.Scan(&title)

	return &Task{Id: fake_id, Title: title}
}

func Find(slice interface{}, f func(value interface{}) bool) int {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return index
			}
		}
	}
	return -1
}
