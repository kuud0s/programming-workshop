package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const filename = "bookmarks.txt"

var bookmarks = map[string]string{}

func saveBookmarks() {
	var content string
	for name, url := range bookmarks {
		content += name + "|" + url + "\n"
	}
	_ = os.WriteFile(filename, []byte(content), 0644)
}

func loadBookmarks() {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return
	}
	data, _ := os.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) == 2 {
			bookmarks[parts[0]] = parts[1]
		}
	}
}

func main() {
	loadBookmarks()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	switch action {
	case "add":
		addBookmark(r)
	case "delete":
		deleteBookmark(r)
	case "exit":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h3>Приложение успешно завершило работу и остановлено.</h3>")
		os.Exit(0)
	}
	renderMenu(w, r)
}

func renderMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Управление закладками</h1>")

	searchQuery := strings.ToLower(r.FormValue("search"))
	searchForm := `
	<form action="/" method="get">
		<table border="0" cellspacing="5" cellpadding="0">
			<tr>
				<td><input type="text" name="search" placeholder="Поиск по названию..." value="%s" size="30"></td>
				<td><button type="submit">Найти</button></td>
				<td><button type="button" onclick="location.href='/'">Сбросить</button></td>
			</tr>
		</table>
	</form>`
	fmt.Fprintf(w, searchForm, r.FormValue("search"))

	if len(bookmarks) == 0 {
		fmt.Fprint(w, "<p><em>Список закладок пуст</em></p>")
	} else {
		fmt.Fprint(w, "<h3>Ваши закладки:</h3>")
		fmt.Fprint(w, "<table border='0' cellspacing='5' cellpadding='0'>")

		hasResults := false
		for name, url := range bookmarks {
			if searchQuery != "" && !strings.Contains(strings.ToLower(name), searchQuery) {
				continue
			}
			hasResults = true
			fmt.Fprintf(w, `
			<tr>
				<td>• <a href="%s" target="_blank"><strong>%s</strong></a></td>
				<td>
					<form action="/" method="get" style="display:inline;">
						<input type="hidden" name="action" value="delete">
						<input type="hidden" name="name" value="%s">
						<button type="submit">Удалить</button>
					</form>
				</td>
			</tr>`, url, name, name)
		}

		if !hasResults && searchQuery != "" {
			fmt.Fprint(w, "<tr><td colspan='2'><em>Ничего не найдено</em></td></tr>")
		}
		fmt.Fprint(w, "</table>")
	}

	forms := `
	<h3>Добавить новую закладку:</h3>
	<form action="/" method="get">
		<input type="hidden" name="action" value="add">
		<table border="0" cellspacing="5" cellpadding="0">
			<tr>
			<td>Название:</td>
			<td><input type="text" name="name" placeholder="Google" size="30" required></td>
			</tr>
			<tr>
			<td>Адрес URL:</td>
			<td><input type="text" name="url" placeholder="https://google.com" size="30" required></td>
			</tr>
			<tr>
			<td></td>
			<td><button type="submit">Сохранить закладку</button></td>
			</tr>
		</table>
	</form>
	<br>
	<form action="/" method="get">
		<input type="hidden" name="action" value="exit">
		<button type="submit">Выйти из приложения</button>
	</form>`
	fmt.Fprint(w, forms)
}

func addBookmark(r *http.Request) {
	name := strings.TrimSpace(r.FormValue("name"))
	url := strings.TrimSpace(r.FormValue("url"))
	if name != "" && url != "" {
		bookmarks[name] = url
		saveBookmarks()
	}
}

func deleteBookmark(r *http.Request) {
	name := r.FormValue("name")
	delete(bookmarks, name)
	saveBookmarks()
}
