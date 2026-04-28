package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Todo 任务结构
type Todo struct {
	ID        int
	Text      string
	Done      bool
	CreatedAt time.Time
}

// TodoList 任务列表
type TodoList struct {
	todos []Todo
	nextID int
}

// NewTodoList 创建任务列表
func NewTodoList() *TodoList {
	return &TodoList{
		todos: []Todo{},
		nextID: 1,
	}
}

// Add 添加任务
func (tl *TodoList) Add(text string) int {
	todo := Todo{
		ID:        tl.nextID,
		Text:      text,
		Done:      false,
		CreatedAt: time.Now(),
	}
	tl.todos = append(tl.todos, todo)
	tl.nextID++
	return todo.ID
}

// Complete 完成任务
func (tl *TodoList) Complete(id int) bool {
	for i := range tl.todos {
		if tl.todos[i].ID == id {
			tl.todos[i].Done = true
			return true
		}
	}
	return false
}

// Delete 删除任务
func (tl *TodoList) Delete(id int) bool {
	for i := range tl.todos {
		if tl.todos[i].ID == id {
			tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
			return true
		}
	}
	return false
}

// List 列出所有任务
func (tl *TodoList) List() []Todo {
	return tl.todos
}

// PrintTodoList 打印任务列表
func PrintTodoList(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("暂无任务！")
		return
	}

	fmt.Println("\n=== 任务列表 ===")
	for _, todo := range todos {
		status := " "
		if todo.Done {
			status = "✓"
		}
		fmt.Printf("[%s] ID: %d - %s (创建于: %s)\n", 
			status, todo.ID, todo.Text, todo.CreatedAt.Format("2006-01-02 15:04"))
	}
	fmt.Println("================")
}

func main() {
	fmt.Println("==================================")
	fmt.Println("  Todo List - Go 语言示例项目")
	fmt.Println("==================================")

	todoList := NewTodoList()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n请选择操作:")
		fmt.Println("1. 查看所有任务")
		fmt.Println("2. 添加任务")
		fmt.Println("3. 完成任务")
		fmt.Println("4. 删除任务")
		fmt.Println("5. 退出")
		fmt.Print("请输入选项 (1-5): ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			PrintTodoList(todoList.List())

		case "2":
			fmt.Print("请输入任务内容: ")
			scanner.Scan()
			text := strings.TrimSpace(scanner.Text())
			if text == "" {
				fmt.Println("任务内容不能为空！")
				continue
			}
			id := todoList.Add(text)
			fmt.Printf("任务已添加！ID: %d\n", id)

		case "3":
			fmt.Print("请输入要完成的任务 ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("无效的 ID！")
				continue
			}
			if todoList.Complete(id) {
				fmt.Println("任务已标记为完成！")
			} else {
				fmt.Println("未找到该任务！")
			}

		case "4":
			fmt.Print("请输入要删除的任务 ID: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("无效的 ID！")
				continue
			}
			if todoList.Delete(id) {
				fmt.Println("任务已删除！")
			} else {
				fmt.Println("未找到该任务！")
			}

		case "5":
			fmt.Println("再见！")
			return

		default:
			fmt.Println("无效的选项，请重新输入！")
		}
	}
}
