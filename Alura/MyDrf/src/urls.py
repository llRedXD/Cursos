# Django imports
from django.contrib import admin
from django.urls import path, include, re_path


urlpatterns = [
    path('admin/', admin.site.urls),
    re_path(r'^v1/', include(('src.escola.urls', 'v1'), namespace='v1')),
    re_path(r'^v2/', include(('src.escola.urls', 'v2'), namespace='v2')),
]
