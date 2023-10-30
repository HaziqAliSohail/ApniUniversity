// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"ApniUniversity/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name ApniUniversity --spec ../../swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.ApniUniversityAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ApniUniversityAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.IntegerProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("integer producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()

	if api.AddAccountHandler == nil {
		api.AddAccountHandler = operations.AddAccountHandlerFunc(func(params operations.AddAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddAccount has not yet been implemented")
		})
	}
	if api.AddClassHandler == nil {
		api.AddClassHandler = operations.AddClassHandlerFunc(func(params operations.AddClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddClass has not yet been implemented")
		})
	}
	if api.AddOrRemoveStudentHandler == nil {
		api.AddOrRemoveStudentHandler = operations.AddOrRemoveStudentHandlerFunc(func(params operations.AddOrRemoveStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddOrRemoveStudent has not yet been implemented")
		})
	}
	if api.AddStudentHandler == nil {
		api.AddStudentHandler = operations.AddStudentHandlerFunc(func(params operations.AddStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddStudent has not yet been implemented")
		})
	}
	if api.AddSubjectHandler == nil {
		api.AddSubjectHandler = operations.AddSubjectHandlerFunc(func(params operations.AddSubjectParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddSubject has not yet been implemented")
		})
	}
	if api.AddTeacherHandler == nil {
		api.AddTeacherHandler = operations.AddTeacherHandlerFunc(func(params operations.AddTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddTeacher has not yet been implemented")
		})
	}
	if api.AssignClassHandler == nil {
		api.AssignClassHandler = operations.AssignClassHandlerFunc(func(params operations.AssignClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AssignClass has not yet been implemented")
		})
	}
	if api.AssignSubjectToStudentHandler == nil {
		api.AssignSubjectToStudentHandler = operations.AssignSubjectToStudentHandlerFunc(func(params operations.AssignSubjectToStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AssignSubjectToStudent has not yet been implemented")
		})
	}
	if api.AssignSubjectToTeacherHandler == nil {
		api.AssignSubjectToTeacherHandler = operations.AssignSubjectToTeacherHandlerFunc(func(params operations.AssignSubjectToTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AssignSubjectToTeacher has not yet been implemented")
		})
	}
	if api.DeleteAccountHandler == nil {
		api.DeleteAccountHandler = operations.DeleteAccountHandlerFunc(func(params operations.DeleteAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteAccount has not yet been implemented")
		})
	}
	if api.DeleteClassHandler == nil {
		api.DeleteClassHandler = operations.DeleteClassHandlerFunc(func(params operations.DeleteClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteClass has not yet been implemented")
		})
	}
	if api.DeleteStudentHandler == nil {
		api.DeleteStudentHandler = operations.DeleteStudentHandlerFunc(func(params operations.DeleteStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteStudent has not yet been implemented")
		})
	}
	if api.DeleteSubjectHandler == nil {
		api.DeleteSubjectHandler = operations.DeleteSubjectHandlerFunc(func(params operations.DeleteSubjectParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteSubject has not yet been implemented")
		})
	}
	if api.DeleteTeacherHandler == nil {
		api.DeleteTeacherHandler = operations.DeleteTeacherHandlerFunc(func(params operations.DeleteTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteTeacher has not yet been implemented")
		})
	}
	if api.GetAccountByIDHandler == nil {
		api.GetAccountByIDHandler = operations.GetAccountByIDHandlerFunc(func(params operations.GetAccountByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetAccountByID has not yet been implemented")
		})
	}
	if api.GetAccountsHandler == nil {
		api.GetAccountsHandler = operations.GetAccountsHandlerFunc(func(params operations.GetAccountsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetAccounts has not yet been implemented")
		})
	}
	if api.GetClassHandler == nil {
		api.GetClassHandler = operations.GetClassHandlerFunc(func(params operations.GetClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetClass has not yet been implemented")
		})
	}
	if api.GetClassByIDHandler == nil {
		api.GetClassByIDHandler = operations.GetClassByIDHandlerFunc(func(params operations.GetClassByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetClassByID has not yet been implemented")
		})
	}
	if api.GetClassOfTeacherHandler == nil {
		api.GetClassOfTeacherHandler = operations.GetClassOfTeacherHandlerFunc(func(params operations.GetClassOfTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetClassOfTeacher has not yet been implemented")
		})
	}
	if api.GetClassesOfStudentHandler == nil {
		api.GetClassesOfStudentHandler = operations.GetClassesOfStudentHandlerFunc(func(params operations.GetClassesOfStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetClassesOfStudent has not yet been implemented")
		})
	}
	if api.GetDefaultedStudentAccountsHandler == nil {
		api.GetDefaultedStudentAccountsHandler = operations.GetDefaultedStudentAccountsHandlerFunc(func(params operations.GetDefaultedStudentAccountsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetDefaultedStudentAccounts has not yet been implemented")
		})
	}
	if api.GetSStudentsOfClassHandler == nil {
		api.GetSStudentsOfClassHandler = operations.GetSStudentsOfClassHandlerFunc(func(params operations.GetSStudentsOfClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSStudentsOfClass has not yet been implemented")
		})
	}
	if api.GetStudentHandler == nil {
		api.GetStudentHandler = operations.GetStudentHandlerFunc(func(params operations.GetStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetStudent has not yet been implemented")
		})
	}
	if api.GetStudentAccountsHandler == nil {
		api.GetStudentAccountsHandler = operations.GetStudentAccountsHandlerFunc(func(params operations.GetStudentAccountsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetStudentAccounts has not yet been implemented")
		})
	}
	if api.GetStudentByIDHandler == nil {
		api.GetStudentByIDHandler = operations.GetStudentByIDHandlerFunc(func(params operations.GetStudentByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetStudentByID has not yet been implemented")
		})
	}
	if api.GetStudentsOfSubjectHandler == nil {
		api.GetStudentsOfSubjectHandler = operations.GetStudentsOfSubjectHandlerFunc(func(params operations.GetStudentsOfSubjectParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetStudentsOfSubject has not yet been implemented")
		})
	}
	if api.GetStudentsOfTeacherHandler == nil {
		api.GetStudentsOfTeacherHandler = operations.GetStudentsOfTeacherHandlerFunc(func(params operations.GetStudentsOfTeacherParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetStudentsOfTeacher has not yet been implemented")
		})
	}
	if api.GetSubjectHandler == nil {
		api.GetSubjectHandler = operations.GetSubjectHandlerFunc(func(params operations.GetSubjectParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSubject has not yet been implemented")
		})
	}
	if api.GetSubjectByIDHandler == nil {
		api.GetSubjectByIDHandler = operations.GetSubjectByIDHandlerFunc(func(params operations.GetSubjectByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSubjectByID has not yet been implemented")
		})
	}
	if api.GetSubjectsOfClassHandler == nil {
		api.GetSubjectsOfClassHandler = operations.GetSubjectsOfClassHandlerFunc(func(params operations.GetSubjectsOfClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSubjectsOfClass has not yet been implemented")
		})
	}
	if api.GetSubjectsOfStudentHandler == nil {
		api.GetSubjectsOfStudentHandler = operations.GetSubjectsOfStudentHandlerFunc(func(params operations.GetSubjectsOfStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetSubjectsOfStudent has not yet been implemented")
		})
	}
	if api.GetTeacherAccountsHandler == nil {
		api.GetTeacherAccountsHandler = operations.GetTeacherAccountsHandlerFunc(func(params operations.GetTeacherAccountsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTeacherAccounts has not yet been implemented")
		})
	}
	if api.GetTeacherByIDHandler == nil {
		api.GetTeacherByIDHandler = operations.GetTeacherByIDHandlerFunc(func(params operations.GetTeacherByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTeacherByID has not yet been implemented")
		})
	}
	if api.GetTeachersHandler == nil {
		api.GetTeachersHandler = operations.GetTeachersHandlerFunc(func(params operations.GetTeachersParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTeachers has not yet been implemented")
		})
	}
	if api.GetTeachersOfClassHandler == nil {
		api.GetTeachersOfClassHandler = operations.GetTeachersOfClassHandlerFunc(func(params operations.GetTeachersOfClassParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTeachersOfClass has not yet been implemented")
		})
	}
	if api.GetTeachersOfSubjectHandler == nil {
		api.GetTeachersOfSubjectHandler = operations.GetTeachersOfSubjectHandlerFunc(func(params operations.GetTeachersOfSubjectParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTeachersOfSubject has not yet been implemented")
		})
	}
	if api.UpdateAccountHandler == nil {
		api.UpdateAccountHandler = operations.UpdateAccountHandlerFunc(func(params operations.UpdateAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateAccount has not yet been implemented")
		})
	}
	if api.UpdateClassNameHandler == nil {
		api.UpdateClassNameHandler = operations.UpdateClassNameHandlerFunc(func(params operations.UpdateClassNameParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateClassName has not yet been implemented")
		})
	}
	if api.UpdateGPAHandler == nil {
		api.UpdateGPAHandler = operations.UpdateGPAHandlerFunc(func(params operations.UpdateGPAParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateGPA has not yet been implemented")
		})
	}
	if api.UpdateStudentNameHandler == nil {
		api.UpdateStudentNameHandler = operations.UpdateStudentNameHandlerFunc(func(params operations.UpdateStudentNameParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateStudentName has not yet been implemented")
		})
	}
	if api.UpdateSubjectNameHandler == nil {
		api.UpdateSubjectNameHandler = operations.UpdateSubjectNameHandlerFunc(func(params operations.UpdateSubjectNameParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateSubjectName has not yet been implemented")
		})
	}
	if api.UpdateTeacherNameHandler == nil {
		api.UpdateTeacherNameHandler = operations.UpdateTeacherNameHandlerFunc(func(params operations.UpdateTeacherNameParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateTeacherName has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
