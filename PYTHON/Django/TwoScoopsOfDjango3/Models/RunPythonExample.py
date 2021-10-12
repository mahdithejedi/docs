from django.db import migrations, models
from django.db.models import manager
from django.db.models.fields import proxy


def forward_func(apps, schema_editor):
    # We get the model from the versioned app registry;
    # if we directly import it, it'll be the wrong version
    Country = apps.get_model("myapp", "Country")
    db_alias = schema_editor.connection.alias
    Country.objects.using(db_alias).bulk_create([
        Country(name="USA", code="us"),
        Country(name="France", code="fr"),
    ])



class Migration(migrations.gration):
    dependencies = [
        # firs care about dependencies
        ('SeprateLogic', '0002_auto_20210905_1355')
    ]
    operations = [
            migrations.RunPython(
                # you can use noop for unreversable codes(idempotent)
                # [for exmaple when we want to migrate two tables to each other which reverse of it is impossible]
                # in this case if you rollback your migrations it wont effect
                forward_func, migrations.RunPython.noop
            )
    ]
