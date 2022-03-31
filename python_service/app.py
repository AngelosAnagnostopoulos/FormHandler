import io
import numpy as np

from matplotlib import pyplot as plt
from matplotlib.backends.backend_agg import FigureCanvasAgg as FigureCanvas
from flask import Response, render_template

from models import Person
from init import create_app
from database import get_all

app = create_app()

a = [0, 0]
l = ["Yes", "No"]


@app.route('/', methods=['GET'])
def fetch():
    persons = get_all(Person)
    all_persons = ()
    for person in persons:
        if person.phone:
            a[0] += 1
        else:
            a[1] += 1
    return render_template("index.html")


@app.route('/plot.png')
def plot_png():
    fig = create_figure()
    output = io.BytesIO()
    FigureCanvas(fig).print_png(output)
    return Response(output.getvalue(), mimetype='image/png')


def create_figure():
    fig, ax = plt.subplots(figsize=(5, 5))
    plt.plot(a, l, ".")
    plt.ylabel('Given')
    plt.xlabel('# of phones')
    plt.title(r'Histogram of Phones given: Y, N')
    return fig
