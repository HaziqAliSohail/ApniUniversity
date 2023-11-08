package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/HaziqAliSohail/ApniUniversity"
	"github.com/HaziqAliSohail/ApniUniversity/gen/restapi/operations"
)

type Handler *operations.ApniUniversityAPI

func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewApniUniversityAPI(spec)

	//Account Handlers
	handler.AddAccountHandler = NewAddAccount(rt)
	handler.GetAccountsHandler = NewGetAccounts(rt)
	handler.GetAccountByIDHandler = NewGetAccountByID(rt)
	handler.UpdateAccountHandler = NewUpdateAccount(rt)
	handler.DeleteAccountHandler = NewDeleteAccount(rt)
	handler.GetStudentAccountsHandler = NewGetStudentAccounts(rt)
	handler.GetDefaultedStudentAccountsHandler = NewGetDefaultedStudentAccounts(rt)
	handler.GetTeacherAccountsHandler = NewGetTeacherAccounts(rt)

	//Class Handlers
	handler.AddClassHandler = NewAddClass(rt)
	handler.GetClassesHandler = NewGetClasses(rt)
	handler.GetClassByIDHandler = NewGetClassByID(rt)
	handler.UpdateClassNameHandler = NewUpdateClassName(rt)
	handler.AddOrRemoveStudentHandler = NewAddOrRemoveStudent(rt)
	handler.GetStudentsOfClassHandler = NewGetStudentsOfClass(rt)
	handler.GetSubjectsOfClassHandler = NewGetSubjectsOfClass(rt)
	handler.GetTeachersOfClassHandler = NewGetTeachersOfClass(rt)
	handler.DeleteClassHandler = NewDeleteClass(rt)

	//Student Handlers
	handler.AddStudentHandler = NewAddStudent(rt)
	handler.AssignSubjectToStudentHandler = NewAssignSubjectToStudent(rt)
	handler.GetStudentsHandler = NewGetStudents(rt)
	handler.GetStudentByIDHandler = NewGetStudentByID(rt)
	handler.GetClassesOfStudentHandler = NewGetClassesOfStudent(rt)
	handler.GetSubjectsOfStudentHandler = NewGetSubjectsOfStudent(rt)
	handler.UpdateStudentNameHandler = NewUpdateStudentName(rt)
	handler.UpdateGPAHandler = NewUpdateGPA(rt)
	handler.DeleteStudentHandler = NewDeleteStudent(rt)

	//Subject Handlers
	handler.AddSubjectHandler = NewAddSubject(rt)
	handler.AssignClassHandler = NewAssignClass(rt)
	handler.GetSubjectsHandler = NewGetSubjects(rt)
	handler.GetSubjectByIDHandler = NewGetSubjectByID(rt)
	handler.GetStudentsOfSubjectHandler = NewGetStudentsOfSubject(rt)
	handler.GetTeacherOfSubjectHandler = NewGetTeacherOfSubject(rt)
	handler.UpdateSubjectNameHandler = NewUpdateSubjectName(rt)
	handler.DeleteSubjectHandler = NewDeleteSubject(rt)

	//Teacher Handlers
	handler.AddTeacherHandler = NewAddTeacher(rt)
	handler.AssignSubjectToTeacherHandler = NewAssignSubjectToTeacher(rt)
	handler.GetTeachersHandler = NewGetTeachers(rt)
	handler.GetTeacherByIDHandler = NewGetTeacherByID(rt)
	handler.GetStudentsOfTeacherHandler = NewGetStudentsOfTeacher(rt)
	handler.GetClassOfTeacherHandler = NewGetClassOfTeacher(rt)
	handler.UpdateTeacherNameHandler = NewUpdateTeacherName(rt)
	handler.DeleteTeacherHandler = NewDeleteTeacher(rt)

	return handler
}
