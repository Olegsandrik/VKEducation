from django.db import models
from django.db.models import Count


class QuestionLikeManager(models.Manager):
    def toggle_like(self, user, question):
        if self.filter(author=user, question=question).exists():
            self.filter(author=user, question=question).delete()
            question.likes -= 1
        else:
            self.create(author=user, question=question)
            question.likes += 1
        question.save()

    def get_likes_for_question(self, question):
        return self.filter(question=question)

    def get_all_likes(self):
        print(self.all())


class AnswerLikeManager(models.Manager):

    def toggle_like(self, user, answer):
        if self.filter(author=user, answer=answer).exists():
            self.filter(author=user, answer=answer).delete()
            answer.likes -= 1
        else:
            self.create(author=user, answer=answer)
            answer.likes += 1
        answer.save()

    def get_likes_for_answer(self, answer):
        return self.filter(answer=answer)

    def get_all_likes(self):
        print(self.all())


class TagManager(models.Manager):
    def create_tag(self, text):
        tag = self.model(text=text)
        tag.save()
        return tag

    def get_top_tag(self):
        return self.annotate(question_count=Count('questions')).order_by('-question_count')[:5]


class AnswerManager(models.Manager):
    def create_answer(self, text, user, question, likes, datetime):
        answer = self.model(context=text, author=user, question=question, likes=likes, datetime=datetime)
        answer.save()
        return answer
    def sort_by_likes(self):
        return self.annotate(num_likes=Count('likesanswer')).order_by('-num_likes')


class QuestionManager(models.Manager):
    def create_question(self, text, user, likes, datetime, tags):
        answer = self.model(context=text, author=user, likes=likes, datetime=datetime, tags=tags)
        answer.save()
        return answer

    def get_questions_with_tag(self, tag_name):
        return self.filter(tags__tag_name=tag_name)


    def get_hot_questions(self):
        return (self.annotate(num_answers=Count('answers')).order_by('-num_answers'))

    def get_new_questions(self):
        return self.order_by('-datetime')

    def get_all_questions(self):
        return self.filter().all().count()