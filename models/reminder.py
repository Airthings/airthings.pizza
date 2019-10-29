from peewee import DateTimeField, ForeignKeyField, Model, AutoField

from .base_model import BaseModel
from .invitation import Invitation


class Reminder(BaseModel):
    id = AutoField(primary_key=True)
    invitation = ForeignKeyField(Invitation, backref="reminders")
    sent_at = DateTimeField()
