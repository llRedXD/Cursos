# Generated by Django 4.0.5 on 2022-06-06 13:41

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('meta_class', '0004_alter_cidade_table_alter_pessoa_table'),
    ]

    operations = [
        migrations.AlterModelOptions(
            name='pessoa',
            options={'ordering': ['-idade', 'nome'], 'verbose_name': 'Pessoa', 'verbose_name_plural': 'Pessoas'},
        ),
        migrations.AlterUniqueTogether(
            name='pessoa',
            unique_together={('nome', 'idade')},
        ),
    ]
