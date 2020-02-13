package template

type Todo struct {
	Title string
	Done  bool
}

type TodoPageDate struct {
	PageTitle string
	Todos     []Todo
}

// return a html template
func TempTodoPage() TodoPageDate {
	data := TodoPageDate{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	return data
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}
