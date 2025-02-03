from django.contrib.auth.models import User
from django.core.management.base import BaseCommand
from django.db import transaction
from faker import Faker
from random import choice, random

from app.models import Question, Answer, Tag, Profile, LikeAnswer, LikeQuestion

class Command(BaseCommand):
    help = 'Заполнение базы данных тестовыми данными.'

    def add_arguments(self, parser):
        parser.add_argument('ratio', type=int, help='Множитель для создания данных')

    @transaction.atomic
    def handle(self, *args, **kwargs):
        ratio = kwargs['ratio']
        fake = Faker()

        profiles = []
        tags = []
        questions = []
        answers = []

        for i in range(ratio):
            username = f"{fake.first_name()}{i}"
            password = fake.password()
            email = fake.email()
            user = User.objects.create_user(username=username, password=password, email=email)
            profile = Profile.objects.create(user=user, premium=False)
            profiles.append(profile)
            profiles.append(profile)

            tag_name = f"{fake.unique.word}{i}"
            tag = Tag.objects.create(tag_name=tag_name)
            tags.append(tag)

        for _ in range(ratio * 10):
            title = fake.sentence(nb_words=6)
            hashed_title = hash(title)
            content = fake.paragraph(nb_sentences=3)
            author = choice(profiles)
            num_tags_to_add = min(len(tags), 3)
            question = Question.objects.create(title=title, context=content, author=author)
            tag_indices = {abs(hashed_title + i) % len(tags) for i in range(num_tags_to_add)}
            question_tags = [tags[i] for i in tag_indices]
            question.tags.set(question_tags)
            questions.append(question)

        for _ in range(ratio * 100):
            content = fake.paragraph(nb_sentences=2)
            author = choice(profiles)
            question = choice(questions)
            answer = Answer.objects.create(context=content, author=author, question=question)
            answers.append(answer)

        for _ in range(ratio * 100):
            user = choice(profiles)
            if choice([True, False]):
                question = choice(questions)
                LikeQuestion.objects.create(author=user, question=question)
            else:
                answer = choice(answers)
                LikeAnswer.objects.create(author=user, answer=answer)

        self.stdout.write(self.style.SUCCESS('Мусор в БД'))
