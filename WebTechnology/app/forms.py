from django import forms
from django.contrib.auth.models import User
from django.core.exceptions import ValidationError

from app.models import Profile, Answer, Question


class LoginForm(forms.Form):
    username = forms.CharField()
    password = forms.CharField(min_length=3, widget=forms.PasswordInput) # textarea

    class Meta:
        model = User
        fields = ('username', 'email', 'password')

    def clean_username(self):
        username = self.cleaned_data.get('username')
        if not (User.objects.filter(username=username).exists()):
            self.add_error('username', 'This username does not exist.')
            raise forms.ValidationError("")
        return username



class RegisterForm(forms.ModelForm):
    password = forms.CharField(widget=forms.PasswordInput)
    password_check = forms.CharField(widget=forms.PasswordInput)

    class Meta:
        model = User
        fields = ('username', 'email', 'password')

    def clean_username(self):
        username = self.cleaned_data.get('username')
        if User.objects.filter(username=username).exists():
            self.add_error('username', 'This username is already in use. Please use a different username.')
            raise forms.ValidationError("")
        return username

    def clan_pass_equal(self):
        pass1 = self.cleaned_data['password']
        pass2 = self.cleaned_data['password_check']
        if pass1 != pass2:
            self.add_error('password', 'Password and password_chek must be equal')
            self.add_error('password_check', '')
            raise forms.ValidationError('')
        return pass1


    def clean_pass_len(self):
        pass1 = self.cleaned_data['password']
        if len(pass1) <= 4:
            self.add_error('password', 'Password must be more than 4 characters.')
            self.add_error('password_check', '')
            raise forms.ValidationError('')
        return pass1

    def clean(self):
        self.clan_pass_equal()
        self.clean_pass_len()
        self.clean_username()
        cleaned_data = super().clean()
        return cleaned_data

    def save(self):
        self.cleaned_data.pop('password_check')
        if 'username' in self.cleaned_data:
            user = User.objects.create_user(**self.cleaned_data)
            Profile.objects.create(user=user, premium=False)
            return user
        else:
            raise forms.ValidationError('error in form')


class EditProfileForm(forms.ModelForm):
    password = forms.CharField(widget=forms.PasswordInput)
    class Meta:
        model = User
        fields = ('username', 'email', 'password')

    def clean_username(self):
        username = self.cleaned_data.get('username')
        if User.objects.filter(username=username).exists():
            self.add_error('username', 'This username is already in use. Please change on a different username.')
            raise forms.ValidationError("")
        return username

    def clean_password(self):
        password = self.cleaned_data['password']
        if len(password) <= 3:
            self.add_error('password', 'Password must be more than 4 characters. ')
            raise forms.ValidationError("")
        return password

    def clean(self):
        #self.clean_password()
        self.clean_username()
        cleaned_data = super().clean()
        return cleaned_data


class AddQuestionForm(forms.ModelForm):
    class Meta:
        model = Question
        fields = ('title', 'context', 'tags')

    def clean_title(self):
        title = self.cleaned_data.get('title')
        if Question.objects.filter(title=title).exists():
            self.add_error('title', 'This question is already in ask. Please change question.')
            raise forms.ValidationError("")
        return title

    def create_qustion(self, Profile):
        question = Question.objects.create(author=Profile, title=self.cleaned_data['title'],
                                           context=self.cleaned_data['context'],)
        question.tags.set(self.cleaned_data['tags'])
        return question


class AddAnswerForm(forms.ModelForm):
    class Meta:
        model = Answer
        fields = ('title', 'context',)

    def clean_title(self):
        title = self.cleaned_data.get('title')
        if len(title)<=4:
            self.add_error('title', 'Title must be more than 4 characters. ')
            raise forms.ValidationError("")
        return title

    def create_answer(self, Profile, Question):
        answer = Answer.objects.create(author=Profile, question=Question,title=self.cleaned_data['title'],
                                           context=self.cleaned_data['context'],)
        return answer