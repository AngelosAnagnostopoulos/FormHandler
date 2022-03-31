import flask_sqlalchemy

db = flask_sqlalchemy.SQLAlchemy()


class Person(db.Model):
    __tablename__ = 'Person'
    id = db.Column(db.Integer, primary_key=True)
    fname = db.Column(db.String())
    email = db.Column(db.String())
    phone = db.Column(db.String())
    city = db.Column(db.String())
