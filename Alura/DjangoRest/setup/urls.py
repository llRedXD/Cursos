from django.contrib import admin
from django.db import router
from django.urls import path, include 

from escola.views import AlunosViewSet, CursosViewSet, MatriculaViewSet, ListaMatriculaAluno, ListaAlunoCurso
from rest_framework import routers

router = routers.DefaultRouter()
router.register('alunos', AlunosViewSet, basename='Alunos')
router.register('cursos', CursosViewSet, basename='Cursos')
router.register('matriculas', MatriculaViewSet, basename='Matricula')

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include(router.urls)),
    path('alunos/<int:pk>/matriculas/', ListaMatriculaAluno.as_view()),
    path('cursos/<int:pk>/matriculas/', ListaAlunoCurso.as_view()),
]
