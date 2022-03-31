from flask import Flask
import flask_sqlalchemy

from models import db
from config import DATABASE_CONNECTION_URI


def create_app():
    flask_app = Flask(__name__, static_url_path='/static')
    flask_app.config['SQLALCHEMY_DATABASE_URI'] = DATABASE_CONNECTION_URI
    flask_app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
    flask_app.app_context().push()
    db.init_app(flask_app)
    db.create_all()
    return flask_app
