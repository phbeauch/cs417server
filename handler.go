package main

import(
    "fmt"
    "os"
    "strconv"
    "encoding/json"  
    "net/http" 
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/gorilla/mux"
)

func getSession() *mgo.Session{
    s, err:=mgo.Dial("mongodb://localhost")
    if err!= nil{
        fmt.Printf("FAILED TO CONNECT TO MONGODB\n")
        os.Exit(1)
    }   
    s.SetSafe(&mgo.Safe{})
    return s
}
//got header information from https://golang.org/pkg/net/http/#pkg-examples
func AddStudent(w http.ResponseWriter,r *http.Request){
    session:=getSession()
    student:=Student{}
    json.NewDecoder(r.Body).Decode(&student)
    student.Rating="D"
    session.DB("studentinfo").C("students").Insert(student)
    result, _:=json.Marshal(student)
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w,"%s",result)    
}

func GetStudent(w http.ResponseWriter,r *http.Request){
    session:=getSession() 
    vars:=mux.Vars(r)
    name:=vars["name"]
    student:=Student{}
    if err:=session.DB("studentinfo").C("students").Find(bson.M{"name": name}).One(&student); err!=nil{
        w.WriteHeader(404)
        fmt.Fprintf(w,"Error! Cannot find: %s\n", name)
        return
    } 
    result, _:=json.Marshal(student)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w,"%s",result)
}

func ListStudents(w http.ResponseWriter,r *http.Request){
    session:=getSession()
    var Students []Student
    if err:=session.DB("studentinfo").C("students").Find(nil).All(&Students); err !=nil{
        w.WriteHeader(404)
        fmt.Fprintf(w,"Error! Cannot ListStudents.\n")
        return
    }
    result, _:=json.Marshal(Students)
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w,"%s",result)
}

func UpdateStudent(w http.ResponseWriter,r *http.Request){
    session:=getSession()
    var Students []Student
    average:=0
    rating:=""
    if err:=session.DB("studentinfo").C("students").Find(nil).All(&Students); err !=nil{
        w.WriteHeader(404)
        fmt.Fprintf(w,"Error! Cannot Update.\n")
        return
    }
    for i:=0;i<len(Students);i++ {
        average+=Students[i].Grade
    }
    average=average/len(Students)
    for i:=0;i<len(Students);i++{
        switch{
        case Students[i].Grade>average+10:
            rating="A"
        case average-10<Students[i].Grade && Students[i].Grade<=average+10:
            rating="B"
        case average-20<Students[i].Grade && Students[i].Grade<=average-10:
            rating="C"
        case Students[i].Grade<=average-20:
            continue
        }   
        session.DB("studentinfo").C("students").UpdateId(Students[i].NetID, bson.M{"$set": bson.M{"rating": rating}})    
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w,"Students Updated!\n")
}

func DeleteStudent(w http.ResponseWriter, r *http.Request){
    session:=getSession()
    var Students []Student
    vars:=mux.Vars(r)
    year, _:=strconv.Atoi(vars["year"])
    if err:=session.DB("studentinfo").C("students").Find(nil).All(&Students); err !=nil{
        w.WriteHeader(404)
        fmt.Fprintf(w,"Error! Cannot Delete!\n")
        return
    }
    for i:=0;i<len(Students);i++{
        if Students[i].Year<year {
            session.DB("studentinfo").C("students").RemoveId(Students[i].NetID)
        }
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w,"Students Deleted!\n")
}
