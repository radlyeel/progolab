package main 

import "fmt" 
import "html/template"
import "net/http"

type Rsvp struct { 
    Name, Email, Phone string 
    WillAttend bool 
}

// A slice (variable sized array) of Rsvp pointers 
// with capacity 10 and initial size 0
var responses = make([]*Rsvp, 0, 10)
// map[keytype]valuetype; see import html/template above
var templates = make(map[string]*template.Template, 3) 

func loadTemplates() { 
    // TODO - load templates here 
    templateNames := [5]string { "welcome", "form", "thanks", "sorry", "list" } 
    for index, name := range templateNames { 
        t, err := template.ParseFiles("layout.html", name + ".html") 
        if (err == nil) { 
            templates[name] = t 
            fmt.Println("Loaded template", index, name) 
        } else { 
            panic(err) 
        } 
    }
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) { 
    templates["welcome"].Execute(writer, nil) 
} 

func listHandler(writer http.ResponseWriter, request *http.Request) { 
    templates["list"].Execute(writer, responses) 
}

type formData struct { 
    *Rsvp 
    Errors []string 
}

func formHandler(writer http.ResponseWriter, request *http.Request) { 
    if request.Method == http.MethodGet { 
        templates["form"].Execute(writer, formData { 
            Rsvp: &Rsvp{}, Errors: []string {}, 
        }) 
    } else if request.Method == http.MethodPost { 
        request.ParseForm() 
        responseData := Rsvp { 
            Name: request.Form["name"][0], 
            Email: request.Form["email"][0], 
            Phone: request.Form["phone"][0], 
            WillAttend: request.Form["willattend"][0] == "true", 
        } 
        errors := []string {} 
        if responseData.Name == "" { 
            errors = append(errors, "Please enter your name") 
        } 
        if responseData.Email == "" { 
            errors = append(errors, "Please enter your email address") 
        } 
        if responseData.Phone == "" { 
            errors = append(errors, "Please enter your phone number") 
        } 
        if len(errors) > 0 { 
            templates["form"].Execute(writer, formData { 
                Rsvp: &responseData, Errors: errors, }) 
            } else {

                responses = append(responses, &responseData) 
                if responseData.WillAttend { 
                    templates["thanks"].Execute(writer, responseData.Name) 
                } else { 
                    templates["sorry"].Execute(writer, responseData.Name) 
                }
            }

        } 
    }

func main() { 
    loadTemplates()

    // Register handlers
    http.HandleFunc("/", welcomeHandler) 
    http.HandleFunc("/list", listHandler)
    http.HandleFunc("/form", formHandler)

    // port 5000 is in use on m2 mac
    // Just specifying the port results in an IPV6 listener on *:7700
    err := http.ListenAndServe("127.0.0.1:7700", nil) 
    if (err != nil) { 
        fmt.Println(err) 
    }
}


