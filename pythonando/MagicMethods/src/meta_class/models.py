from tabnanny import verbose
from django.db import models
import uuid

class Cidade(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    nome = models.CharField(max_length=255)
    idade = models.IntegerField()

    class Meta:
        db_table = "contact_city"

    def __str__(self):
        return self.nome

class Pessoa(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    nome = models.CharField(max_length=255)
    idade = models.IntegerField()
    cidade = models.ForeignKey(Cidade, on_delete=models.CASCADE)

    class Meta:
        ordering = ['-idade', 'nome']
        db_table = "contact_user"
        verbose_name = 'Pessoa'
        verbose_name_plural = 'Pessoas'
        unique_together = ['nome', 'idade']

    def __str__(self):
        return f"Nome {self.nome} | Idade {self.idade}"