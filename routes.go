package main

import (
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route
var routes=Routes{
    Route{
        "AddStudent",
        "POST",
        "/Student",
        AddStudent,
    },
    Route{
        "GetStudent",
        "GET",
        "/Student/getstudent",
        GetStudent,
    },
    Route{
        "ListStudents",
        "GET",
        "/Student/listall",
        ListStudents,
    },
    Route{
        "UpdateStudent",
        "PUT",
        "/Student",
        UpdateStudent,     
    },
    Route{
        "DeleteStudent",
        "DELETE",
        "/Student/{year:[0-9]+}",
        DeleteStudent,
    },
}
