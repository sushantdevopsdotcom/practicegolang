// package routes

// import (
//     "github.com/gin-gonic/gin"
//     "big-web-app/handlers"
//     "big-web-app/middleware"
// )

// func SetupRouter() *gin.Engine {
//     r := gin.Default()
//     r.Use(middleware.BasicAuthMiddleware()) // Apply basic auth middleware
//     r.GET("/items", handlers.GetItems)
//     r.POST("/items", handlers.CreateItem)
//     return r
// }


package routes

import (
    "github.com/gin-gonic/gin"
    "big-web-app/handlers"
    "big-web-app/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(middleware.BasicAuthMiddleware()) // Apply basic auth middleware
    r.GET("/items", handlers.GetItems)
    r.GET("/item-headers", handlers.GetItemHeaders) // Add route for item headers
    r.POST("/items", handlers.CreateItem)
    return r
}
