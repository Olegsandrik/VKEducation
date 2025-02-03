from django.contrib.auth.decorators import login_required
from django.shortcuts import render, redirect
from django.http import HttpResponse, Http404
from django.core.paginator import Paginator
from django.contrib.auth import login, authenticate, logout
from django.urls import reverse
from django.views.decorators.csrf import csrf_protect

from app.forms import LoginForm, RegisterForm, EditProfileForm, AddQuestionForm, AddAnswerForm
from app.models import Profile, Question, Answer


@csrf_protect
def ask(request):
    reg = []
    notreg = []
    if request.user.is_authenticated:
        reg.append(1)
        if request.method == 'GET':
            ask_form = AddQuestionForm()
        if request.method == 'POST':
            ask_form = AddQuestionForm(request.POST)
            if ask_form.is_valid():
                question = ask_form.create_qustion(request.user.profile)
                if question:
                    print('sucsess saving')
                    return redirect(reverse('question', kwargs={'number': question.id}))
                else:
                    ask_form.add_error(None, "Error saving question")
    else:
        notreg.append(1)
        if request.method == 'GET':
            ask_form = AddQuestionForm()
        if request.method == 'POST':
            ask_form = AddQuestionForm(request.POST)
            if ask_form.is_valid():
                ask_form.add_error(None, "You are not logged in")
            else:
                ask_form.add_error(None, "Invalid data")
    return render(request, "app/ask.html", {'user': request.user,
                                                 'register': reg, 'notregister': notreg,
                                            'form': ask_form})

@csrf_protect
def mylogout(request):
    logout(request)
    return redirect(request.GET.get('next', '/'))


def index(request):
    questionsnew = Question.objects.get_new_questions()
    set_page = []
    for i in range(1, min(len(questionsnew), 6), 1):
        set_page.append(f"{i}")
    if len(questionsnew)==1:
        set_page = ['1']
    page = request.GET.get("page", '1')
    if page in set_page:
        return render(request, "app/test.html", {'questions': paginate(questionsnew, page, 3),
                                                'pages': set_page})
    raise Http404('Страница не найдена')


def hot(request):
    questionshot = Question.objects.get_hot_questions()
    set_page = []
    for i in range(1, min(len(questionshot), 6), 1):
        set_page.append(f"{i}")
    if len(questionshot) == 1:
            set_page = ['1']
    page = request.GET.get("page", '1')
    if page in set_page:
        reg = []
        notreg = []
        if request.user.is_authenticated:
            reg.append(1)
        else:
            notreg.append(1)
        return render(request, "app/hot.html", {'questions': paginate(questionshot, page, 3),
                                                 'pages': set_page, 'user': request.user,
                                                 'register': reg, 'notregister': notreg})
    raise Http404('Страница не найдена')


def home(request):
    questionsnew = Question.objects.get_new_questions()
    set_page = []
    for i in range(1, min(len(questionsnew), 4), 1):
        set_page.append(f"{i}")
    if len(questionsnew)==1:
        set_page = ['1']
    page = request.GET.get("page", '1')
    if page in set_page:
        reg = []
        notreg = []
        if request.user.is_authenticated:
            reg.append(1)
        else:
            notreg.append(1)
        return render(request, "app/home.html", {'questions': paginate(questionsnew, page, 3),
                                                 'pages': set_page, 'user': request.user,
                                                 'register': reg, 'notregister': notreg})
    raise Http404('Страница не найдена')


@csrf_protect
def mylogin(request):
    if request.method == 'GET':
        login_form = LoginForm()
    if request.method == 'POST':
        login_form = LoginForm(request.POST)
        if login_form.is_valid():
            user = authenticate(request, **login_form.cleaned_data)
            #print(user)
            if user is not None:
                login(request, user)
                #print("sucsess")
                return redirect(request.GET.get('next', '/')) # request.Get.get('continue', '/') или revers('home')
            else:
                login_form.add_error('password', "Wrong password")
    return render(request, "app/login.html", context={'form': login_form})


@csrf_protect
def singup(request):
    if request.method == 'GET':
        registr_form = RegisterForm()
    if request.method == 'POST':
        registr_form = RegisterForm(request.POST)
        if registr_form.is_valid():
            user = registr_form.save()
            if user:
                print("sucsessfully registered")
                login(request, user)
                return redirect(reverse('home')) # request.Get.get('continue', '/') или revers('home')
            else:
                registr_form.add_error(None, "User saving error!")
    return render(request, "app/singup.html", context={'form': registr_form})


def tag(request, tagname):
    questionstag = Question.objects.get_questions_with_tag(tagname)
    set_page = []
    for i in range(1, min(len(questionstag), 6), 1):
        set_page.append(f"{i}")
    if len(questionstag)==1:
        set_page = ['1']
    page = request.GET.get("page", '1')
    if page in set_page:
        reg = []
        notreg = []
        if request.user.is_authenticated:
            reg.append(1)
        else:
            notreg.append(1)
        return render(request, "app/tag.html", {'questions': paginate(questionstag, page, 3),
                                                 'pages': set_page, 'tagname': tagname, 'user': request.user,
                                                 'register': reg, 'notregister': notreg})
    raise Http404('Страница не найдена')



def question(request, number):
    questionsreal = Question.objects.all()
    set_page = []
    for i in range(1, min(len(questionsreal), 4), 1):
        set_page.append(f"{i}")
    page = request.GET.get("page", '1')
    item = questionsreal[number-1]
    if page in set_page:
        reg = []
        notreg = []
        ans_form = AddAnswerForm()
        if request.user.is_authenticated:
            reg.append(1)
        else:
            notreg.append(1)
        return render(request, "app/question.html", {'question': item,
                                                     'pages': set_page, 'user': request.user,
                                                 'register': reg, 'notregister': notreg,
                                                     'form': ans_form})
    raise Http404('Страница не найдена')


def paginate(objects_list, page, per_page=10):
    paginator = Paginator(objects_list, per_page)
    return paginator.page(page)

@csrf_protect
def edit(request):
    reg = []
    notreg = []
    if request.user.is_authenticated:
        reg.append(1)
        if request.method == 'GET':
            edit_form = EditProfileForm()
        if request.method == 'POST':
            edit_form = EditProfileForm(request.POST)
            if edit_form.is_valid():
                user = request.user
                user.username = edit_form.cleaned_data.get('username')
                user.email = edit_form.cleaned_data.get('email')
                user.set_password(edit_form.cleaned_data['password'])
                user.save()
                login(request, user)
                print('sucsess edit')
            else:
                edit_form.add_error(None, 'Invalid data')
        return render(request, "app/edit.html", {'user': request.user,
                                                 'register': reg, 'notregister': notreg, 'form': edit_form})
    else:
        notreg.append(1)
        return Http404

