from django.apps import AppConfig


class MetaClassConfig(AppConfig):
    default_auto_field = 'django.db.models.BigAutoField'
    name = 'meta_class'
