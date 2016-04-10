package character

import (
	"html/template"
	"net/http"
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
	if err := characterTemplate.Execute(w, characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var characterTemplate = template.Must(template.New("roster").Parse(`
<html>
  <head>
    <title>CS496 Assignment 2 - loughlid</title>
  </head>
  <body>
    <div><h1>Cast of Characters Database</h1></div>
    <div><h2>Current Roster</h2>
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
          <th>Date Added</th>
        </tr>
      </thead>
      <tbody>
        {{range .}}
        <tr>
          <td>{{.Id}}</td>
          <td>{{.LastName}}</td>
          <td>{{.FirstName}}</td>
          <td>{{.Sex}}</td>
          <td>{{.Vegan}}</td>
          <td>{{.Email}}</td>
          <td>{{.PhoneNum}}</td>
          <td>{{.Date}}</td>
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
