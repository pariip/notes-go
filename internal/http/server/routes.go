package server

func (s *httpServer) setRoutes() {
	//public routes
	{
		s.public.POST("/login", s.handler.login)
		s.public.POST("/signup", s.handler.signup)
	}

	//admin routes
	{
		s.admin.POST("/user", s.handler.adminCreateUser)
		s.admin.GET("/user/:id", s.handler.adminGetUser)
		s.admin.PATCH("/user", s.handler.adminUpdateUser)
		s.admin.DELETE("/user/:id", s.handler.adminDeleteUser)
	}

	//user routes
	{
		s.user.GET("/token/refresh/:id", s.handler.refreshToken)
	}

	//note routes
	{
		s.note.POST("", s.handler.createNote)
		s.note.GET("", s.handler.getAllNotes)
		s.note.GET("/:id", s.handler.getNoteByID)
		s.note.PATCH("", s.handler.updateNote)
		s.note.DELETE("/:id", s.handler.deleteNote)
	}
}
