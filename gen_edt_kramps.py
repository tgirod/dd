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
"""

import random

# table du personnel
per = ["AG-"+str(mat) for mat in random.choices(range(1,25), k=10)]
# table des prisonniers
cap = ["PA-21" ]+ ["PA-"+str(mat) for mat in random.choices( range(1,20), k=3)] + ["PA-"+str(mat) for mat in random.choices( range(30,50), k=2)]
cap += ["PF-47"] + ["PF-"+str(mat) for mat in random.choices( range(1,45), k=4)]


#print( "per", per)
#print( "cap", cap)

# table des lieux
loc = ["AC1", "AC2", "AC3", "DZ", "RR", "CE"]

# table des heures
hh = ["10", "14", "16"]

def form_register( key, val, desc, priv ):
    reg = '{"'+key+', '+str(val)+', "'+desc+'", '+str(priv)+'},'
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
#print( "**** personnel", per_edt )
for k,v in per_edt.items():
    print( form_register( k, v, "", 1))
for k,v in cap_edt.items():
    print( form_register( k, v, "", 1))
print( len(per_edt)+len(cap_edt) )
