
package main 

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type album struct {

    ID      string   `json:id`
    Title   string   `json:title`
    Artist  string   `json:artist`
    Price  float64  `json:priece`

}

var albums = []album{
    
    { ID : "1",Title:"The Night",Artist:"Caleb Seña",Price:50.34 },
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},

}

// getAlbums responds with the list of all albums as JSON.

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    
    // 3 part of the tutorial 
    router.GET("/albums/:id", getAlbumByID )


    // 2 Part of the tutorial 
    router.POST("/albums",postAlbums)

    router.Run("localhost:8080")
}

// Part of the tutorial 

// postAlbums adds an album  from JSON recived int the request body



/* 
From a different command line window, use curl to make a request to your running web service.

$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'


   
For See in the console 
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
*/


func postAlbums  (c * gin.Context){
    var newAlbum album 

    // Call  bindJSOM to bind  the recibed JSON to New ALbum 

    if err := c.BindJSON(&newAlbum); err != nil {
        return 
    }
    
    // add the new album to the slice .
    albums = append(albums , newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum )
}


// getAlbumByID  loscates the album whose ID value matches the id 
// parameter sent by the client , then returns that album  as a reponse


func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}



/*
From a different command line window, use curl to make a request to your running web service.

$ curl http://localhost:8080/albums/2

*/
































