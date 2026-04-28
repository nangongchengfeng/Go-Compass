# Issue 1.5: Todo List 命令行项目 - 学习知识点

## 📚 概述

这是我们的第一个完整项目。我们将用前面学到的知识，构建一个实用的 Todo List 命令行程序。这个项目涵盖：结构体、切片、函数、命令行交互等多个知识点。

---

## 📋 需求分析与设计

### 1.1 功能需求

我们的 Todo List 需要实现以下功能：
1. 添加任务
2. 查看所有任务
3. 标记任务完成
4. 删除任务
5. 命令行交互界面

### 1.2 数据结构设计

我们需要定义一个 Todo 结构体来表示任务：
- ID：任务编号
- Content：任务内容
- Done：是否完成
- CreatedAt：创建时间

---

## 📝 实现步骤

### 2.1 定义 Todo 结构体

首先，我们需要定义 Todo 结构体来存储任务信息：

```go
package main

import (
    "fmt"
    "time"
)

// Todo 结构体定义
type Todo struct {
    ID        int
    Content   string
    Done      bool
    CreatedAt time.Time
}
```

### 2.2 创建 TodoList 管理器

我们需要一个结构体来管理所有的 Todo 任务：

```go
// TodoList 管理器
type TodoList struct {
    todos []Todo
    nextID int
}

// NewTodoList 创建一个新的 TodoList
func NewTodoList() *TodoList {
    return &TodoList{
        todos: make([]Todo, 0),
        nextID: 1,
    }
}
```

---

## 🎯 功能实现

### 3.1 添加任务功能

实现添加任务的方法：

```go
// AddTodo 添加新任务
func (tl *TodoList) AddTodo(content string) {
    todo := Todo{
        ID:        tl.nextID,
        Content:   content,
        Done:      false,
        CreatedAt: time.Now(),
    }
    tl.todos = append(tl.todos, todo)
    tl.nextID++
    fmt.Printf("✅ 任务已添加！ID: %d\n", todo.ID)
}
```

### 3.2 查看任务列表功能

实现查看所有任务的方法：

```go
// ListTodos 查看所有任务
func (tl *TodoList) ListTodos() {
    if len(tl.todos) == 0 {
        fmt.Println("📭 暂无任务")
        return
    }
    
    fmt.Println("\n📋 任务列表:")
    fmt.Println("----------------------------------------")
    for _, todo := range tl.todos {
        status := "⬜"
        if todo.Done {
            status = "✅"
        }
        fmt.Printf("%s ID: %d | %s\n", status, todo.ID, todo.Content)
        fmt.Printf("   创建时间: %s\n", todo.CreatedAt.Format("2006-01-02 15:04:05"))
        fmt.Println("----------------------------------------")
    }
}
```

### 3.3 完成任务功能

实现标记任务完成的方法：

```go
// MarkDone 标记任务完成
func (tl *TodoList) MarkDone(id int) bool {
    for i := range tl.todos {
        if tl.todos[i].ID == id {
            tl.todos[i].Done = true
            fmt.Printf("✅ 任务 %d 已标记为完成！\n", id)
            return true
        }
    }
    fmt.Printf("❌ 未找到 ID 为 %d 的任务\n", id)
    return false
}
```

### 3.4 删除任务功能

实现删除任务的方法：

```go
// DeleteTodo 删除任务
func (tl *TodoList) DeleteTodo(id int) bool {
    for i := range tl.todos {
        if tl.todos[i].ID == id {
            tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
            fmt.Printf("🗑️  任务 %d 已删除！\n", id)
            return true
        }
    }
    fmt.Printf("❌ 未找到 ID 为 %d 的任务\n", id)
    return false
}
```

---

## 💻 命令行交互界面

### 4.1 显示菜单

实现一个简单的菜单显示函数：

```go
// ShowMenu 显示菜单
func ShowMenu() {
    fmt.Println("\n========== Todo List ==========")
    fmt.Println("1. 添加任务")
    fmt.Println("2. 查看任务列表")
    fmt.Println("3. 标记任务完成")
    fmt.Println("4. 删除任务")
    fmt.Println("5. 退出")
    fmt.Println("================================")
    fmt.Print("请选择操作 (1-5): ")
}
```

### 4.2 主程序流程

实现完整的主程序：

```go
package main

import (
    "fmt"
    "time"
    "bufio"
    "os"
    "strconv"
    "strings"
)

// Todo 结构体定义
type Todo struct {
    ID        int
    Content   string
    Done      bool
    CreatedAt time.Time
}

// TodoList 管理器
type TodoList struct {
    todos []Todo
    nextID int
}

// NewTodoList 创建一个新的 TodoList
func NewTodoList() *TodoList {
    return &TodoList{
        todos: make([]Todo, 0),
        nextID: 1,
    }
}

// AddTodo 添加新任务
func (tl *TodoList) AddTodo(content string) {
    todo := Todo{
        ID:        tl.nextID,
        Content:   content,
        Done:      false,
        CreatedAt: time.Now(),
    }
    tl.todos = append(tl.todos, todo)
    tl.nextID++
    fmt.Printf("✅ 任务已添加！ID: %d\n", todo.ID)
}

// ListTodos 查看所有任务
func (tl *TodoList) ListTodos() {
    if len(tl.todos) == 0 {
        fmt.Println("📭 暂无任务")
        return
    }
    
    fmt.Println("\n📋 任务列表:")
    fmt.Println("----------------------------------------")
    for _, todo := range tl.todos {
        status := "⬜"
        if todo.Done {
            status = "✅"
        }
        fmt.Printf("%s ID: %d | %s\n", status, todo.ID, todo.Content)
        fmt.Printf("   创建时间: %s\n", todo.CreatedAt.Format("2006-01-02 15:04:05"))
        fmt.Println("----------------------------------------")
    }
}

// MarkDone 标记任务完成
func (tl *TodoList) MarkDone(id int) bool {
    for i := range tl.todos {
        if tl.todos[i].ID == id {
            tl.todos[i].Done = true
            fmt.Printf("✅ 任务 %d 已标记为完成！\n", id)
            return true
        }
    }
    fmt.Printf("❌ 未找到 ID 为 %d 的任务\n", id)
    return false
}

// DeleteTodo 删除任务
func (tl *TodoList) DeleteTodo(id int) bool {
    for i := range tl.todos {
        if tl.todos[i].ID == id {
            tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
            fmt.Printf("🗑️  任务 %d 已删除！\n", id)
            return true
        }
    }
    fmt.Printf("❌ 未找到 ID 为 %d 的任务\n", id)
    return false
}

// ShowMenu 显示菜单
func ShowMenu() {
    fmt.Println("\n========== Todo List ==========")
    fmt.Println("1. 添加任务")
    fmt.Println("2. 查看任务列表")
    fmt.Println("3. 标记任务完成")
    fmt.Println("4. 删除任务")
    fmt.Println("5. 退出")
    fmt.Println("================================")
    fmt.Print("请选择操作 (1-5): ")
}

func main() {
    tl := NewTodoList()
    scanner := bufio.NewScanner(os.Stdin)
    
    for {
        ShowMenu()
        
        scanner.Scan()
        choice := scanner.Text()
        
        switch choice {
        case "1":
            fmt.Print("请输入任务内容: ")
            scanner.Scan()
            content := scanner.Text()
            if strings.TrimSpace(content) != "" {
                tl.AddTodo(content)
            } else {
                fmt.Println("❌ 任务内容不能为空！")
            }
            
        case "2":
            tl.ListTodos()
            
        case "3":
            fmt.Print("请输入要标记完成的任务 ID: ")
            scanner.Scan()
            idStr := scanner.Text()
            id, err := strconv.Atoi(idStr)
            if err == nil {
                tl.MarkDone(id)
            } else {
                fmt.Println("❌ 请输入有效的数字！")
            }
            
        case "4":
            fmt.Print("请输入要删除的任务 ID: ")
            scanner.Scan()
            idStr := scanner.Text()
            id, err := strconv.Atoi(idStr)
            if err == nil {
                tl.DeleteTodo(id)
            } else {
                fmt.Println("❌ 请输入有效的数字！")
            }
            
        case "5":
            fmt.Println("👋 再见！")
            return
            
        default:
            fmt.Println("❌ 无效的选择，请重新输入！")
        }
    }
}
```

---

## 🚀 运行程序

### 5.1 编译和运行

把上面的代码保存到 `main.go` 文件中，然后运行：

```bash
go run main.go
```

或者编译成可执行文件：

```bash
go build -o todo main.go
./todo
```

### 5.2 使用示例

```
========== Todo List ==========
1. 添加任务
2. 查看任务列表
3. 标记任务完成
4. 删除任务
5. 退出
================================
请选择操作 (1-5): 1
请输入任务内容: 学习 Go 语言复合数据类型
✅ 任务已添加！ID: 1

========== Todo List ==========
1. 添加任务
2. 查看任务列表
3. 标记任务完成
4. 删除任务
5. 退出
================================
请选择操作 (1-5): 1
请输入任务内容: 完成 Todo List 项目
✅ 任务已添加！ID: 2

========== Todo List ==========
1. 添加任务
2. 查看任务列表
3. 标记任务完成
4. 删除任务
5. 退出
================================
请选择操作 (1-5): 2

📋 任务列表:
----------------------------------------
⬜ ID: 1 | 学习 Go 语言复合数据类型
   创建时间: 2024-01-15 14:30:00
----------------------------------------
⬜ ID: 2 | 完成 Todo List 项目
   创建时间: 2024-01-15 14:31:00
----------------------------------------
```

---

## 📝 练习题

### 基础题（5 道）

1. **修改程序**：添加一个功能，显示已完成和未完成任务的数量统计。
2. **添加功能**：实现一个"清空所有任务"的功能。
3. **修改界面**：给菜单添加更多的装饰，让界面更好看。
4. **错误处理**：当用户输入无效的 ID 时，给出更友好的提示。
5. **代码优化**：将菜单显示部分的代码提取成单独的函数。

### 中级题（5 道）

6. **搜索功能**：添加根据关键词搜索任务的功能。
7. **任务分类**：给 Todo 添加一个 Category 字段，可以给任务分类（工作、学习、生活等）。
8. **优先级**：添加 Priority 字段（高、中、低），并在显示时按优先级排序。
9. **截止日期**：添加 DueDate 字段，可以设置任务的截止日期。
10. **数据持久化**：使用 JSON 文件保存任务数据，下次启动时自动加载。

### 高级题（5 道）

11. **编辑功能**：添加编辑任务内容的功能。
12. **子任务**：每个任务可以有子任务。
13. **标签系统**：给任务添加标签功能，一个任务可以有多个标签。
14. **统计报表**：添加统计功能，显示每周完成任务数量等。
15. **Web 界面**：使用简单的 HTTP 服务器，给 Todo List 添加 Web 界面。

---

## ✅ 验收标准

### 基础语法（3 项）
- [ ] 能够正确定义和使用结构体
- [ ] 能够正确使用切片存储和管理数据
- [ ] 能够理解和实现方法（method）

### 功能实现（5 项）
- [ ] 能够独立实现添加任务功能
- [ ] 能够独立实现查看任务列表功能
- [ ] 能够独立实现完成任务功能
- [ ] 能够独立实现删除任务功能
- [ ] 能够实现简单的命令行交互界面

### 代码质量（3 项）
- [ ] 代码结构清晰，函数职责单一
- [ ] 有适当的注释说明
- [ ] 代码能够正常编译运行

### 扩展能力（3 项）
- [ ] 能够完成至少 3 道基础练习题
- [ ] 能够完成至少 2 道中级练习题
- [ ] 能够尝试完成 1 道高级练习题

---

## 📚 参考资料

- Go 官方文档：https://go.dev/doc/
- Go by Example：https://gobyexample.com/
- The Go Programming Language（书籍）

---

完成了这个项目，你已经迈出了 Go 语言学习的重要一步。继续加油！
