package main

import (
	"adsTool/containers"
	"adsTool/pkg/builder"
	"adsTool/pkg/googleads"
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	// Mở database SQLite
	var err error
	db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	googleApi := googleads.NewGoogleApisClient()
	defer googleApi.Close()

	//createTable()

	home := builder.NewWindow(&containers.Home{})
	home.ShowAndRun()
}

func openAddNewWindow(app fyne.App) fyne.Window {
	addWindow := app.NewWindow("Thêm hành đông chuyển đổi")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter name")

	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter email")

	submitButton := widget.NewButton("Submit", func() {
		name := nameEntry.Text
		email := emailEntry.Text
		if name != "" && email != "" {
			_, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
			if err != nil {
				dialog.ShowError(err, addWindow)
			} else {
				addWindow.Close()
			}
		}
	})

	form := container.NewVBox(
		widget.NewForm(
			widget.NewFormItem("Name", nameEntry),
			widget.NewFormItem("Email", emailEntry),
		),
		submitButton,
	)

	addWindow.SetContent(form)
	addWindow.Resize(fyne.NewSize(400, 200))
	return addWindow
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func getRowCount() int {
	rows, err := db.Query("SELECT COUNT(*) FROM users")
	if err != nil {
		return 0
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		rows.Scan(&count)
	}
	return count
}

func getName(row int) string {
	rows, err := db.Query("SELECT name FROM users LIMIT 1 OFFSET ?", row)
	if err != nil {
		return ""
	}
	defer rows.Close()
	var name string
	if rows.Next() {
		rows.Scan(&name)
	}
	return name
}

func getEmail(row int) string {
	rows, err := db.Query("SELECT email FROM users LIMIT 1 OFFSET ?", row)
	if err != nil {
		return ""
	}
	defer rows.Close()
	var email string
	if rows.Next() {
		rows.Scan(&email)
	}
	return email
}
