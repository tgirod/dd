#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
Génération de l'emploi du temps des prisonniers de la Kramps sous la forme

PT_ID_LOC_HH = False/True
CT_ID_LOC_HH = False/True
PT: PersonnelTimetable
CT: CaptiveTimetable
ID: id du personnel ou du prisonnier (ex: PF-47 pour Sasquatch)
LOC:lieu/salle (ex: AC2 atelier démontage 2)
HH: heure, par créneau de 2h => 10,14,16

ATTENTION : Installer Faker pour générer des noms et des dates
pip install Faker


"""
import faker
from datetime import date

fakeN = faker.Faker(['es_ES', 'de_DE', 'it_IT', 'en_US', 'fr_FR'])
start_date = date(1950,1,1)
end_date = date(2002,12,31)
fakeD = faker.Faker(['fr_FR'])


import random
import string

# ******************************************************************* form_entry
def form_entry( id, entry ):
    ent = '{"'+id+', []string{'
    for k in entry["keys"]:
        ent += '"'+k+'", '
    ent += '}'
    ent += ', '+str(entry["priv"])
    ent += ', "'+entry["owner"]+'"'
    ent += ', "'+entry["title"]+'"'
    ent += ', "'+entry["content"]+'"'
    ent += '},'

    return ent

# ******************************************************************************
# ******************************************************************** Personnel
# ******************************************************************************
# table du personnel id = G-[A-Z][1-300]
nb_agent = 10
per = ["G-"+str(mat)+str(nb) for mat,nb in zip([random.choice(string.ascii_uppercase) for _ in range(nb_agent)],
                                              random.choices( range(1,300), k=nb_agent))]
per += ["G-A37", "G-C3"]


per_dict = {}
for p in per:
    name = fakeN.name()
    date = fakeD.date_between_dates( start_date, end_date)
    misc = fakeD.uuid4()
    msg = name
    msg += " - " + f"{date.day}/{date.month}/{date.year}"
    msg += " - " + str(misc)
    
    per_dict[p] = { "keys": ['agent','gardien'],
                    "priv": 1,
                    "owner": "",
                    "title": name,
                    "content": msg }
per_dict["G-A37"]["title"] = "Jonathan Swift"
per_dict["G-A37"]["content"] = "Jonathan Swift  - 11/4/1984 - 383de3ff-56eb-4745-b472-e046ff8e552e"
per_dict["G-C3"]["title"] = "Harvey Zimmermann"
per_dict["G-C3"]["content"] = "Harvey Zimmermann  - 15/11/1961 - a37dd901-6526-4913-8900-daf9af5f8fab"
per = list(per_dict.keys())
random.shuffle(per)
# print in random order
for k in per: 
    print( form_entry( k, per_dict[k] ))


# ******************************************************************************
# ****************************************************************** Prisonniers
# ******************************************************************************
# table des prisonniers id = P[AF][1-300]
nb_cap = 20
cap = ["P"+str(mat)+str(nb) for mat,nb in zip([random.choice(string.ascii_uppercase) for _ in range(nb_cap)],
                                              random.choices( range(1,300), k=nb_cap))]
cap += ["PA-21", "PF-47"]

cap_dict = {}
for id in cap:
    name = fakeN.name()
    date = fakeD.date_between_dates( start_date, end_date)
    misc = fakeD.uuid4()
    msg = name
    msg += " - " + f"{date.day}/{date.month}/{date.year}"
    msg += " - " + str(misc)
    
    cap_dict[id] = { "keys": ['détenu'],
                    "priv": 1,
                    "owner": "",
                    "title": name,
                    "content": msg }
cap_dict["PA-21"]["title"] = "Pedro Ramirez"
cap_dict["PA-21"]["content"] = "Pedro Ramirez - 8/7/1991 - 383ca3ff-58eb-4745-efff-e046ff8e552e"
cap_dict["PF-47"]["title"] = "Stefan Jasinski"
cap_dict["PF-47"]["content"] = "Stefan Jasinski - 15/3/1988 - c54dd982-6526-cd42-8900-cbe27f5f8fab"
cap = list(cap_dict.keys())
random.shuffle(cap)
# print in random order
for k in cap: 
    print( form_entry( k, cap_dict[k] ))



#print( "per", per)
#print( "cap", cap)

# table des lieux
loc = ["AC1", "AC2", "AC3", "DZ", "RR", "CE"]

# table des heures
hh = ["10", "14", "16"]

def form_register( key, val, desc, priv ):
    reg = '{"'+key+'", '+str(val)+', "'+desc+'", '+str(priv)+'},'
    return reg

# generate edt-pers
def gen_edt( start, id, loc, hh):
    edt = {}
    for p in id:
        for h in hh:
            sel_l = random.choice( loc )
            for l in loc:
                key = start+'_'+p+'_'+l+'_'+h
                val = (l == sel_l)
                edt[key] = val
    return edt

per_edt = gen_edt( 'PT', per, loc, hh)
cap_edt = gen_edt( 'DT', cap, loc, hh)
##print( "**** personnel", per_edt )
for k,v in per_edt.items():
    print( form_register( k, v, "", 1))
for k,v in cap_edt.items():
    print( form_register( k, v, "", 1))

print( len(per_edt)+len(cap_edt) )

    

