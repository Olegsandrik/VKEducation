from django.contrib.auth.models import User
from django.contrib.contenttypes.fields import GenericForeignKey, GenericRelation
from django.contrib.contenttypes.models import ContentType
from django.db import models

from app.manegers import QuestionLikeManager, AnswerLikeManager, TagManager, QuestionManager, AnswerManager


class Question(models.Model):
    title = models.CharField(max_length=256)
    author = models.ForeignKey('Profile', on_delete=models.CASCADE, related_name='questions')
    datetime = models.DateTimeField(auto_now_add=True)
    likes = models.IntegerField(default=0)
    context = models.TextField()
    tags = models.ManyToManyField('Tag', related_name='questions')
    objects = QuestionManager()

    def __str__(self):
        return f"{self.author.user.username} question: {self.title}"


class LikeQuestion(models.Model):
    question = models.ForeignKey('Question', on_delete=models.CASCADE, related_name='likesquestion')
    author = models.ForeignKey('Profile', on_delete=models.CASCADE, related_name='likesquestion')
    objects = QuestionLikeManager()

    def __str__(self):
        return f"{self.author.user.username} like question {self.question.title}"


class LikeAnswer(models.Model):
    answer = models.ForeignKey('Answer', on_delete=models.CASCADE, related_name='likesanswer')
    author = models.ForeignKey('Profile', on_delete=models.CASCADE, related_name='likesanswer')
    objects = AnswerLikeManager()

    def __str__(self):
        return f"{self.author.user.username} like answer {self.answer.title}"


class Answer(models.Model):
    title = models.CharField(max_length=256)
    context = models.TextField()
    author = models.ForeignKey('Profile', on_delete=models.CASCADE, related_name='answers')
    datetime = models.DateTimeField(auto_now_add=True)
    tags = models.ManyToManyField('Tag', related_query_name='answers')
    likes = models.IntegerField(default=0)
    question = models.ForeignKey('Question', on_delete=models.CASCADE, related_name='answers', null=True, blank=False)
    objects = AnswerManager()

    def __str__(self):
        return f"{self.author.user.username} answer: {self.question}"


class Tag(models.Model):
    tag_name = models.CharField(max_length=100, null=True, blank=False)
    objects = TagManager()

    def __str__(self):
        return f"{self.tag_name} - tag"


def user_directory_path(instance, filename):
    return 'avatar_{0}/{1}'.format(instance.user.username, filename)


class Profile(models.Model):
    premium = models.BooleanField(default=False)
    user = models.OneToOneField(User, on_delete=models.CASCADE)
    image = models.ImageField(
        upload_to=user_directory_path,
        blank=True,
        null=True
    )

    def __str__(self):
        return self.user.username


class myuser(models.Model):
    username = models.CharField(max_length=25)
    usersurname = models.CharField(max_length=25)
    email = models.EmailField()
    password = models.CharField (max_length=25)
    datetime = models.DateTimeField(auto_now=True)
