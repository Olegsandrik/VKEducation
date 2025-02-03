from django.urls import path
from . import views
urlpatterns = [
    path('', views.home, name='home'),
    path('ask/', views.ask, name='ask'),
    path('login/', views.mylogin, name='login'),
    path('singup/', views.singup, name='singup'),
    path('hot/', views.hot, name='hot'),
    path('tag/<str:tagname>/', views.tag, name='tag'),
    path('question/<int:number>/', views.question, name='question'),
    path('index/', views.index, name='index'),
    path('mylogout/', views.mylogout, name='mylogout'),
    path('pfofile/edit/', views.edit, name='edit'),
]
