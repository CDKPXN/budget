package main

import (
	"kunpeng/budget-api/config"
	"kunpeng/budget-api/controllers"

	// "kunpeng/budget-api/database"
	"kunpeng/budget-api/middleware"

	"github.com/betacraft/yaag/yaag"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
)

/**
 * 初始化 iris app
 * @method NewApp
 * @return  {[type]}  api      *iris.Application  [iris app]
 */
func newApp() (api *iris.Application) {
	api = iris.New()
	api.Use(logger.New())

	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(controllers.APIResource(false, nil, "404 Not Found"))
	})
	api.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Oups something went wrong, try again")
	})

	// database.Init()

	//"github.com/iris-contrib/middleware/cors"
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	appName := config.Conf.Get("app.name").(string)
	appDoc := config.Conf.Get("app.doc").(string)
	appURL := config.Conf.Get("app.url").(string)
	//api 文档配置
	yaag.Init(&yaag.Config{ // <- IMPORTANT, init the middleware.
		On:       true,
		DocTitle: appName,
		DocPath:  appDoc + "/index.html", //设置绝对路径
		BaseUrls: map[string]string{
			"Production": appURL,
			"Staging":    "",
		},
	})

	v1 := api.Party("/api", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Use(middleware.NewYaag()) // <- IMPORTANT, register the middleware.
		v1.Get("/", func(ctx iris.Context) {
			ctx.JSON("ok")
		})
		v1.PartyFunc("/finance", func(finance router.Party) {
			finance.Get("/", controllers.GetFinance)
			finance.Post("/", controllers.CreateFinance)
			finance.Put("/{id:string}", controllers.UpdateFinance)
		})
		v1.PartyFunc("/staff", func(staff router.Party) {
			staff.Get("/", controllers.GetStaff)
			staff.Post("/", controllers.CreateStaff)
			staff.Put("/{id:string}", controllers.UpdateStaff)
		})
		v1.PartyFunc("/project", func(project router.Party) {
			project.Get("/", controllers.GetProject)
			// project.Get("/{id:string}", controller.GetProjectByID)
			project.Post("/", controllers.CreateProject)
			project.Delete("/{id:string}", controllers.DeleteProject)
			project.Put("/{id:string}", controllers.UpdateProject)

			project.PartyFunc("/{id:string}/workload", func(workload router.Party) {
				workload.Post("/", controllers.CreateProjectWorkload)
				// workload.Put("/{id:string}", controller.UpdateProjectWorkload)
			})
		})
	}

	return
}

func main() {
	app := newApp()

	addr := config.Conf.Get("app.addr").(string)
	app.Run(iris.Addr(addr))
}
