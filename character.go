package character

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"appengine"
	"appengine/datastore"
)

// Attributes of the Character entity
type Character struct {
	Id        int64 `datastore:"-"`
	FirstName string
	LastName  string
	Sex       string
	Vegan     string
	Email     string
	PhoneNum  string
	Date      time.Time
}

// Setup web server
func init() {
	http.HandleFunc("/", root)
	// View only page
	http.HandleFunc("/edit", edit)

	http.HandleFunc("/update", update)
}

// The characterKey returns the key used for all character entries.
func characterKey(c appengine.Context) *datastore.Key {
	// The string "default_character" here could be varied to have multiple rosters.
	return datastore.NewKey(c, "Characters", "default_character", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	q := datastore.NewQuery("Character").Ancestor(characterKey(c)).Order("-Date").Limit(10)
	characters := make([]Character, 0, 10)
	keys, err := q.GetAll(c, &characters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Grab ID from NDB so we can display it
	for i, key := range keys {
		characters[i].Id = key.IntID()
	}
	if err := characterEditTemplate.Execute(w, characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var characterEditTemplate = template.Must(template.New("roster").Parse(`
<html>
  <head>
    <title>CS496 Assignment 2 - loughlid</title>
  </head>
  <body>
    <div><h1>Cast of Characters Database</h1></div>
    <div><h2>Current Roster (edit mode)</h2>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Last Name</th>
          <th>First Name</th>
          <th>Sex</th>
          <th>Vegan</th>
          <th>Email</th>
          <th>Phone Number</th>
          <th>Date</th>
          <th></th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {{range .}}
        <tr>
          <form action="/update" method="get">
            <td><input type="number" name="Id" value="{{.Id}}" hidden required>{{.Id}}</td>
            <td><input type="text" name="LastName" value="{{.LastName}}" required></td>
            <td><input type="text" name="FirstName" value="{{.FirstName}}" required></td>
            {{if eq .Sex "male"}}
              <td><input type="radio" name="Sex" value="male" checked> Male
                <input type="radio" name="Sex" value="female"> Female</td>
            {{end}}
            {{if eq .Sex "female"}}
              <td><input type="radio" name="Sex" value="male"> Male
                <input type="radio" name="Sex" value="female" checked> Female</td>
            {{end}}
            {{if eq .Vegan "Yes"}}
              <td><input type="checkbox" name="Vegan" checked> Yes</td>
            {{end}}
            {{if eq .Vegan "No"}}
              <td><input type="checkbox" name="Vegan"> Yes</td>
            {{end}}
            <td><input type="email" name="Email" value="{{.Email}}" required></td>
            <td><input type="tel" name="PhoneNum" value="{{.PhoneNum}}" required></td>
            <td><input type="text" name="Date" value="{{.Date}}" hidden>{{.Date}}</td>
            <td><input type="submit" name="update_button" value="Update"></td>
            <td><input type="submit" name="delete_button" value="Delete"></td>
          </form>
        </tr>
        {{end}}
      </tbody>
    </table>
    </div>
    <h2>Add a Character To The Roster</h2>
    <form action="/edit" method="post">
      <div>
        <p><label>First Name: <input type="text" name="FirstName" required></label>
        <p><label>Last Name: <input type="text" name="LastName" required></label>
        <p><label>Sex: <input type="radio" name="Sex" value="male" checked> Male 
          <input type="radio" name="Sex" value="female"> Female</label>
        <p><label>Vegan: <input type="checkbox" name="Vegan"> Yes</label>
        <p><label>Email: <input type="email" name="Email" required></label>
        <p><label>Phone number: <input type="tel" name="PhoneNum" required></label>
      </div>
      <div><input type="submit" value="Save"></div>
    </form>
  </body>
</html>
`))

func edit(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	g := Character{
		FirstName: r.FormValue("FirstName"),
		LastName:  r.FormValue("LastName"),
		Sex:       r.FormValue("Sex"),
		Email:     r.FormValue("Email"),
		PhoneNum:  r.FormValue("PhoneNum"),
		Date:      time.Now(),
	}

	// Convert checkbox value to string
	if v := r.FormValue("Vegan"); v == "on" {
		g.Vegan = "Yes"
	} else {
		g.Vegan = "No"
	}
	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Character", characterKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func update(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if r.FormValue("delete_button") == "Delete" {
		Id := r.FormValue("Id")
		s, err := strconv.ParseInt(Id, 10, 64)
		keyId := datastore.NewKey(c, "Character", "", s, characterKey(c))
		query := datastore.NewQuery("Character").Filter("__key__ =", keyId).Limit(10)
		//var people []Character
		people := make([]Character, 0, 10)
		key, err := query.GetAll(c, &people)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = datastore.Delete(c, key[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.FormValue("update_button") == "Update" {
		Id := r.FormValue("Id")
		s, err := strconv.ParseInt(Id, 10, 64)
		keyId := datastore.NewKey(c, "Character", "", s, characterKey(c))
		query := datastore.NewQuery("Character").Filter("__key__ =", keyId).Limit(10)
		//var people []Character
		people := make([]Character, 0, 10)
		key, err := query.GetAll(c, &people)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		g := Character{
			FirstName: r.FormValue("FirstName"),
			LastName:  r.FormValue("LastName"),
			Sex:       r.FormValue("Sex"),
			Email:     r.FormValue("Email"),
			PhoneNum:  r.FormValue("PhoneNum"),
		}
		if v := r.FormValue("Vegan"); v == "on" {
			g.Vegan = "Yes"
		} else {
			g.Vegan = "No"
		}
		_, err = datastore.Put(c, key[0], &g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
