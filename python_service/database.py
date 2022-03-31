from models import db


def get_all(db):
    data = db.query.all()
    return data
