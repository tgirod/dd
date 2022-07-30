import itertools
import random
from itertools import product
import faker

orb = ["LEO", "MEO"]
region = ["EU", "AM", "AS", "AU"]
district = ["D%02d"%(i) for i in range(1,31)]
const = ["SATCOM", "STARLINK", "VIASAT", "IRIDIUM"]
band = ["600MHZ", "800MHZ", "1200MHZ"]

map_region = {
    "EU": "Europole",
    "AM": "Amerique",
    "AS": "Asie",
    "AU": "Australie",
}

f = faker.Faker()

geo = product(["GEO"], region, district, band)
autre = product(orb, const, band)

for sat in list(geo) + list(autre):
    name = "-".join(sat)
    state = "true" if f.pybool() else "false"
    print('{"%s", %s, "%s", 5},'%(name, state, ""))
