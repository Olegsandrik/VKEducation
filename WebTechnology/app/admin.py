from django.contrib import admin
from .models import Question, Answer, Tag, Profile, LikeQuestion, LikeAnswer


admin.site.register(Question)
admin.site.register(Answer)
admin.site.register(Tag)
admin.site.register(Profile)
admin.site.register(LikeAnswer)
admin.site.register(LikeQuestion)
