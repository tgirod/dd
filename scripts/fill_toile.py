#!/usr/bin/env python3
# -*- coding: utf-8 -*-

""" Fill the 'toile' with some Posts
"""

from datetime import datetime, timedelta
from random import choice, randrange    # random integers

users = ["Alain", "Fab", "Tomtom", "Cyn"]
# posts = ["Sujet à traiter en urgence",
#          "Pas méchant mais quand même",
#          "Un dernier verre pour la route"]
posts = ["TRY starduck", "TRY bradburry", "TRY bilongo",
         "TRY Luke Sywalker", "TRY Harry", "Try on y va ?"]


def add_post_thread( topic, title, nb, date ):
    sender = choice(users)
    filename = topic+"/"+filename_from(date, title, sender)
    id_post = 0
    print( f"__OPEN {filename}" )
    with open(filename, 'w') as of:
        of.write( f"Numero {id_post} dans la série")
    while id_post < nb:
        id_post += 1
        sender = choice(users)
        date = next_datetime(date)
        filename = topic+"/"+filename_from(date, "Re: "+title, sender)
        print( f"  answer_{id_post} {filename}" )
        with open(filename, 'w') as of:
            of.write( f"Numero {id_post} dans la série")

def next_datetime( date_time ):
    """ Ajoute entre 1 et 24*60 minutes """
    d_minutes = randrange(1, 24*60)
    time_delta = timedelta( minutes=d_minutes )

    return date_time+time_delta

def filename_from( date_time, title, user ):
    res = date_time.strftime('%y%m%d')
    res += "_"+date_time.strftime('%H%M%S')
    res += "_"+title
    res += "_"+user
    return res

def test_datetime():
    dt = datetime( 2020, 2, 17, 22, 8, 45 )

    print( f"dt = {dt.ctime()}")
    print( f"format d={dt.strftime('%y%m%d')}, i={dt.strftime('%H%M%S')}" )

    dt = next_datetime(dt)
    print( f"format d={dt.strftime('%y%m%d')}, i={dt.strftime('%H%M%S')}" )
    return dt

def fill_toile():
    """Pour chaque 'post' de 'posts"', crée 0-10 réponses"""
    for p in posts:
        add_post_thread( "../toile/dd.local/forum/news",
                         p,
                         randrange(0,11),
                         test_datetime()
                        )


if __name__ == '__main__':
    # add_post_thread( "../forum/news",
    #                  "Un essai de génération",
    #                  10,
    #                  test_datetime()
    #                 )
    fill_toile()
